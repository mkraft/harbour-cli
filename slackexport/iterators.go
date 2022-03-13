package slackexport

import "github.com/mkraft/harbour"

type userIterator struct {
	index int
	users []*harbour.User
}

func (u *userIterator) HasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *userIterator) Next() (*harbour.User, error) {
	if u.HasNext() {
		user := u.users[u.index]
		u.index++
		return user, nil
	}
	return nil, nil
}

type groupIterator struct {
	index  int
	groups []*harbour.Group
}

func (u *groupIterator) HasNext() bool {
	if u.index < len(u.groups) {
		return true
	}
	return false
}

func (u *groupIterator) Next() (*harbour.Group, error) {
	if u.HasNext() {
		group := u.groups[u.index]
		u.index++
		return group, nil
	}
	return nil, nil
}
