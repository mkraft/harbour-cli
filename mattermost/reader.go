package mattermost

import "context"

type Reader interface {
	Users(ctx context.Context) chan *Result[*User]
	Groups(ctx context.Context) chan *Result[*Group]
}

type ResultVal interface {
	*User | *Group
}

type Result[T ResultVal] struct {
	Val T
	Err error
}
