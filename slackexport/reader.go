package slackexport

import "github.com/mkraft/mattermost"

type Reader struct{}

func (s *Reader) Users() chan *mattermost.UserResult {
	results := make(chan *mattermost.UserResult)
	go func() {
		// do some work to retrieve this list of users from the slack export
		for _, user := range []*mattermost.User{
			{Username: "fake.user1"},
			{Username: "fake.user2"},
		} {
			results <- &mattermost.UserResult{User: user}
		}
		close(results)
	}()
	return results
}

func (s *Reader) Groups() chan *mattermost.GroupResult {
	results := make(chan *mattermost.GroupResult)
	go func() {
		// do some work to retrieve this list of groups from the slack export
		for _, group := range []*mattermost.Group{
			{Name: "fake-group1"},
			{Name: "fake-group2"},
		} {
			results <- &mattermost.GroupResult{Group: group}
		}
		close(results)
	}()
	return results
}
