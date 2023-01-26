// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package repo

import (
	"github.com/blackhorseya/irent/pkg/httpx"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

// Injectors from wire.go:

func CreateRepo(opts *Options, httpclient httpx.Client, rw *sqlx.DB) IRepo {
	iRepo := NewImpl(opts, httpclient, rw)
	return iRepo
}

// wire.go:

var testProviderSet = wire.NewSet(NewImpl)
