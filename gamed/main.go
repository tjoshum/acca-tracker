package main

import (
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	"github.com/prometheus/common/log"
	"github.com/tjoshum/acca-tracker/database/proto"
	"github.com/tjoshum/acca-tracker/lib/names"
	"golang.org/x/net/context"
)

func main() {
	cmd.Init()

	cl := database.NewDatabaseServiceClient(names.DatabaseSvc, client.DefaultClient)

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "no-user",
		"X-From-Id": names.GameDaemon,
	})

	rsp, err := cl.AddGame(ctx, &database.AddGameRequest{
		Week:     2,
		HomeTeam: 8,
		AwayTeam: 4,
	})
	if err != nil {
		log.Fatal("AddGame failed", err)
	}
	if rsp.Error != database.ErrorCode_SUCCESS {
		log.Fatal("AddGame failed with response code", rsp.Error.String())
	}
}
