package test

import (
	"account-web/config"
	"account-web/server"
	"testing"
)

func TestServer(t *testing.T) {
	config.FilePath = "../config/config.yaml"
	server.Start()
}
