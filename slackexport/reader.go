package slackexport

import "github.com/mkraft/mattermost"

type Reader struct{}

func (s *Reader) UserIterator() (mattermost.UserIterator, error) {
	// do some work to retrieve this list of users from the slack export
	fakeUserIterator := &userIterator{
		index: 0,
		users: []*mattermost.User{
			{Username: "fake.user1"},
			{Username: "fake.user2"},
		},
	}
	return fakeUserIterator, nil
}

func (s *Reader) GroupIterator() (mattermost.GroupIterator, error) {
	// do some work to retrieve this list of groups from the slack export
	fakeGroupIterator := &groupIterator{
		index: 0,
		groups: []*mattermost.Group{
			{Name: "fake-group1"},
			{Name: "fake-group2"},
		},
	}
	return fakeGroupIterator, nil
}
