package restful

import (
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/blackhorseya/sion/adapter/restful/cmds"
	"github.com/blackhorseya/sion/entity/domain/rental/agg"
	"github.com/blackhorseya/sion/entity/domain/rental/biz"
	"github.com/blackhorseya/sion/entity/domain/rental/model"
	"github.com/blackhorseya/sion/pkg/adapterx"
	"github.com/blackhorseya/sion/pkg/configx"
	"github.com/blackhorseya/sion/pkg/contextx"
	"github.com/blackhorseya/sion/pkg/transports/httpx"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"go.uber.org/zap"
)

type impl struct {
	server *httpx.Server
	bot    *messaging_api.MessagingApiAPI
	rental biz.IRentalBiz

	commands []cmds.TextCommander
}

func newRestful(
	server *httpx.Server,
	bot *messaging_api.MessagingApiAPI,
	rental biz.IRentalBiz,
) adapterx.Restful {
	return &impl{
		server:   server,
		bot:      bot,
		rental:   rental,
		commands: cmds.NewCommands(),
	}
}

func newService(
	server *httpx.Server,
	bot *messaging_api.MessagingApiAPI,
	rental biz.IRentalBiz,
) adapterx.Servicer {
	return newRestful(server, bot, rental)
}

func (i *impl) Start() error {
	ctx := contextx.Background()

	err := i.InitRouting()
	if err != nil {
		return err
	}

	err = i.server.Start(ctx)
	if err != nil {
		return err
	}

	ctx.Info("start restful server", zap.String("addr", configx.C.HTTP.GetAddr()))

	return nil
}

func (i *impl) AwaitSignal() error {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	if sig := <-c; true {
		ctx := contextx.Background()
		ctx.Info("receive signal", zap.String("signal", sig.String()))

		err := i.server.Stop(ctx)
		if err != nil {
			ctx.Error("shutdown restful server error", zap.Error(err))
			return err
		}
	}

	return nil
}

func (i *impl) InitRouting() error {
	api := i.server.Router.Group("/api")
	{
		api.POST("/callback", i.callback)
	}

	return nil
}

func (i *impl) GetRouter() *gin.Engine {
	return i.server.Router
}

//nolint:gocognit // it's okay
func (i *impl) callback(c *gin.Context) {
	ctx, err := contextx.FromGin(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	cb, err := webhook.ParseRequest(configx.C.LineBot.Secret, c.Request)
	if err != nil {
		if errors.Is(err, linebot.ErrInvalidSignature) {
			ctx.Error("invalid line bot signature", zap.Error(err))
			_ = c.Error(err)
		} else {
			ctx.Error("parse line bot request error", zap.Error(err))
			_ = c.Error(err)
		}

		return
	}

	var replyMessages []messaging_api.MessageInterface
	for _, event := range cb.Events {
		switch e := event.(type) {
		case webhook.MessageEvent:
			switch message := e.Message.(type) {
			case webhook.TextMessageContent:
				replyMessages, err = i.handleTextMessage(ctx, message)
				if err != nil {
					ctx.Error("handle text message error", zap.Error(err))
					_ = c.Error(err)
					return
				}

				_, err = i.bot.ReplyMessage(&messaging_api.ReplyMessageRequest{
					ReplyToken: e.ReplyToken,
					Messages:   replyMessages,
				})
				if err != nil {
					ctx.Error("reply message error", zap.Error(err))
					_ = c.Error(err)
					return
				}
			case webhook.LocationMessageContent:
				target := &model.Location{
					Latitude:  message.Latitude,
					Longitude: message.Longitude,
				}

				var assets agg.Assets
				assets, _, err = i.rental.ListByLocation(ctx, target, biz.ListByLocationOptions{
					Page: 1,
					Size: 5,
				})
				if err != nil {
					ctx.Error("list by location error", zap.Error(err))
					_ = c.Error(err)
					return
				}

				container, err2 := assets.FlexContainer()
				if err2 != nil {
					ctx.Error("flex container error", zap.Error(err2))
					_ = c.Error(err2)
					return
				}

				_, err = i.bot.ReplyMessage(&messaging_api.ReplyMessageRequest{
					ReplyToken: e.ReplyToken,
					Messages: []messaging_api.MessageInterface{
						&messaging_api.FlexMessage{
							AltText:  "Available Cars",
							Contents: container,
						},
					},
				})
				if err != nil {
					ctx.Error("reply message error", zap.Error(err))
					_ = c.Error(err)
					return
				}
			default:
				ctx.Debug("message type not support", zap.String("type", e.GetType()))
			}
		default:
			ctx.Debug("event type not support", zap.String("type", e.GetType()))
		}
	}

	c.Status(http.StatusOK)
}

func (i *impl) handleTextMessage(
	ctx contextx.Contextx,
	message webhook.TextMessageContent,
) ([]messaging_api.MessageInterface, error) {
	for _, command := range i.commands {
		messages, err := command.Execute(ctx, message.Text)
		if err != nil {
			return handleError(err)
		}

		if messages != nil {
			return messages, nil
		}
	}

	return nil, errors.New("not support")
}

func handleError(err error) ([]messaging_api.MessageInterface, error) {
	return []messaging_api.MessageInterface{
		messaging_api.TextMessage{Text: err.Error()},
	}, nil
}
