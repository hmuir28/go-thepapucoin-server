package main

import (
	"context"
	"github.com/hmuir28/go-thepapucoin-server/kafka"

	"github.com/hmuir28/go-thepapucoin-server/p2p"
	"github.com/hmuir28/go-thepapucoin-server/database"
)

func main() {

	newInstance := database.NewRedisClient()

	p2pServer := p2p.NewP2PServer()

	var ctx = context.Background()

	go kafka.Subscriber(ctx, p2pServer, newInstance)
	
	p2p.StartServer(ctx, p2pServer, newInstance)

	for{}
}
