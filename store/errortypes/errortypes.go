package errortypes

//NotFound represents a not found error
type NotFound interface {
	error
	ErrNotFound()
}

type notFound struct {
	NotFound
}

//NewNotFound returns a new not found error
func NewNotFound() NotFound {
	return &notFound{}
}
