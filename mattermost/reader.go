package mattermost

type Reader interface {
	Users() chan *Result[*User]
	Groups() chan *Result[*Group]
}

type ResultVal interface {
	*User | *Group
}

type Result[T ResultVal] struct {
	Val T
	Err error
}
