package slackexport

import (
	"context"
	"time"

	"example.com/mattermost"
)

type Reader struct{}

func (s *Reader) Users(ctx context.Context) chan *mattermost.Result[*mattermost.User] {
	results := make(chan *mattermost.Result[*mattermost.User])

	users := []*mattermost.User{
		{Username: "fake.user1"},
		{Username: "fake.user2"},
		{Username: "fake.user3"},
	}

	go func() {
		for _, user := range users {
			select {
			case <-ctx.Done():
				close(results)
				return
			default:
				time.Sleep(2 * time.Second)
				results <- &mattermost.Result[*mattermost.User]{Val: user}
			}
		}
		close(results)
	}()

	return results
}

func (s *Reader) Groups(ctx context.Context) chan *mattermost.Result[*mattermost.Group] {
	results := make(chan *mattermost.Result[*mattermost.Group])

	groups := []*mattermost.Group{
		{Name: "fake-group1"},
		{Name: "fake-group2"},
		{Name: "fake-group3"},
	}

	go func() {
		for _, group := range groups {
			select {
			case <-ctx.Done():
				close(results)
				return
			default:
				time.Sleep(2 * time.Second)
				results <- &mattermost.Result[*mattermost.Group]{Val: group}
			}
		}
		close(results)
	}()

	return results
}
