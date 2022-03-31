package mattermost

type UserResult struct {
	User *User
	Err  error
}

type GroupResult struct {
	Group *Group
	Err   error
}

type Reader interface {
	Users() chan *UserResult
	Groups() chan *GroupResult
}
