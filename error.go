package main

type Error interface {
	error
	Status() int
}

type StatusError struct {
	code int
	Err  error
}

func (s StatusError) Status() int {
	return s.code
}

func (s StatusError) Error() string {
	return s.Err.Error()
}
