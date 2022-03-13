package mattermost

type Writer interface {
	WriteUser(*User) error
	WriteGroup(*Group) error
}
