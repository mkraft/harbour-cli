package mattermost

type Reader interface {
	UserIterator() (UserIterator, error)
	GroupIterator() (GroupIterator, error)
}
