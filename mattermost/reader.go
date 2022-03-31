package mattermost

type Reader interface {
	Users() chan *Result[*User]
	Groups() chan *Result[*Group]
}

type Result[T any] struct {
	Val T
	Err error
}
