package main

import (
	"github.com/lovoo/webgo/internal/api"
	"go.uber.org/zap"
)

func main() {
	s := &api.Server{}
	fatalOnError("server stopped", s.Run())
}

func fatalOnError(msg string, err error) {
	if err != nil {
		zap.L().Fatal(msg, zap.Error(err))
	}
}
