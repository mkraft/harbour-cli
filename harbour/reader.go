package harbour

type Reader interface {
	UserIterator() (UserIterator, error)
	GroupIterator() (GroupIterator, error)
}
