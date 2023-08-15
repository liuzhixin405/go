package iservice

import (
	"context"
	"fmt"
	"multi/model"
)

//排行榜

type TopUserGetter interface {
	GetTopUsers(ctx context.Context, count int) ([]*model.User, error)
}
type TopScoreNotifier interface {
	NotifyTopScore(ctx context.Context, id string, score int) error
}

type Leaderboard struct {
	topUserGetter    TopUserGetter
	topScoreNotifier TopScoreNotifier
}

func New(topUserGetter TopUserGetter, topScoreNotifier TopScoreNotifier) *Leaderboard {
	return &Leaderboard{
		topUserGetter:    topUserGetter,
		topScoreNotifier: topScoreNotifier,
	}
}

func (l *Leaderboard) NotifyTopPlayers(ctx context.Context, top int) error {
	users, err := l.topUserGetter.GetTopUsers(ctx, top)
	if err != nil {
		return fmt.Errorf("topUserGetter.GetTopUsers: %w", err)
	}

	for _, user := range users {
		err = l.topScoreNotifier.NotifyTopScore(ctx, user.ID, user.Score)

		if err != nil {
			return fmt.Errorf("topScoreNotifier.NotifyTopScore: %w", err)
		}
	}
	return nil
}
