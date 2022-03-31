package slackexport

import "github.com/mkraft/mattermost"

type Reader struct{}

func (s *Reader) Users() chan *mattermost.Result[*mattermost.User] {
	results := make(chan *mattermost.Result[*mattermost.User])
	go func() {
		// do some work to retrieve this list of users from the slack export
		for _, user := range []*mattermost.User{
			{Username: "fake.user1"},
			{Username: "fake.user2"},
		} {
			results <- &mattermost.Result[*mattermost.User]{Val: user}
		}
		close(results)
	}()
	return results
}

func (s *Reader) Groups() chan *mattermost.Result[*mattermost.Group] {
	results := make(chan *mattermost.Result[*mattermost.Group])
	go func() {
		// do some work to retrieve this list of groups from the slack export
		for _, group := range []*mattermost.Group{
			{Name: "fake-group1"},
			{Name: "fake-group2"},
		} {
			results <- &mattermost.Result[*mattermost.Group]{Val: group}
		}
		close(results)
	}()
	return results
}
