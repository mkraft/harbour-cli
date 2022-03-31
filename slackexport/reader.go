package slackexport

import (
	"github.com/mkraft/mattermost"
)

type Reader struct{}

func (s *Reader) Users(done chan bool) (chan *mattermost.User, chan error) {
	results := make(chan *mattermost.User)
	errors := make(chan error)
	go func() {
		// do some work to retrieve this list of users from the slack export
		for _, user := range []*mattermost.User{
			{Username: "fake.user1"},
			{Username: "fake.user2"},
		} {
			results <- user
		}
		done <- true
		close(results)
		close(errors)
	}()
	return results, errors
}

func (s *Reader) Groups(done chan bool) (chan *mattermost.Group, chan error) {
	results := make(chan *mattermost.Group)
	errors := make(chan error)
	go func() {
		// do some work to retrieve this list of groups from the slack export
		for _, group := range []*mattermost.Group{
			{Name: "fake-group1"},
			{Name: "fake-group2"},
		} {
			results <- group
		}
		done <- true
		close(results)
		close(errors)
	}()
	return results, errors
}
