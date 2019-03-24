package main

import (
	"github.com/izumin5210/grapi/pkg/grapiserver"
)

func run() error {
	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
		// TODO
		),
	)
	return s.Serve()
}

