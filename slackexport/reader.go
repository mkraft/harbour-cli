package slackexport

import "github.com/mkraft/harbour"

type Reader struct{}

func (s *Reader) UserIterator() (harbour.UserIterator, error) {
	fakeUserIterator := &userIterator{
		index: 0,
		users: []*harbour.User{
			&harbour.User{Username: "fake.user1"},
			&harbour.User{Username: "fake.user2"},
		},
	}
	return fakeUserIterator, nil
}

func (s *Reader) GroupIterator() (harbour.GroupIterator, error) {
	fakeGroupIterator := &groupIterator{
		index: 0,
		groups: []*harbour.Group{
			&harbour.Group{Name: "fake-group1"},
			&harbour.Group{Name: "fake-group2"},
		},
	}
	return fakeGroupIterator, nil
}
