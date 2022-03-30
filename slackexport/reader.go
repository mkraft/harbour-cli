package slackexport

import "github.com/mkraft/mattermost"

type Reader struct{}

func (s *Reader) Users() chan *mattermost.User {
	users := make(chan *mattermost.User)
	go func() {
		// do some work to retrieve this list of users from the slack export
		for _, user := range []*mattermost.User{
			{Username: "fake.user1"},
			{Username: "fake.user2"},
		} {
			users <- user
		}
		close(users)
	}()
	return users
}

func (s *Reader) Groups() chan *mattermost.Group {
	groups := make(chan *mattermost.Group)
	go func() {
		// do some work to retrieve this list of groups from the slack export
		for _, group := range []*mattermost.Group{
			{Name: "fake-group1"},
			{Name: "fake-group2"},
		} {
			groups <- group
		}
		close(groups)
	}()
	return groups
}
