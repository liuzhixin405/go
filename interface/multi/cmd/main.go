package main

import (
	"context"
	"multi/iservice"
	"multi/model"
	"multi/service"
)

func main() {
	database := model.New()
	notifier := service.New()

	leaderboard := iservice.New(database, notifier)

	leaderboard.NotifyTopPlayers(context.Background(), 3)

	iservice.DeleteUserHandler(database)

}
