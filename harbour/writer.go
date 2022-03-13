package harbour

type Writer interface {
	WriteUser(*User) error
	WriteGroup(*Group) error
}
