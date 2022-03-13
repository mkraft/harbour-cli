package harbour

type UserIterator interface {
	HasNext() bool
	Next() (*User, error)
}

type GroupIterator interface {
	HasNext() bool
	Next() (*Group, error)
}
