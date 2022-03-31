package mattermost

type Reader interface {
	Users(done chan bool) (chan *User, chan error)
	Groups(done chan bool) (chan *Group, chan error)
}
