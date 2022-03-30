package mattermost

type Reader interface {
	Users() chan *User
	Groups() chan *Group
}
