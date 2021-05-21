package transport

type responseError struct {
	Err    error
	Msg    string
	Status int
}

func (r responseError) Unwrap() error {
	return r.Err
}

func (r responseError) Error() string {
	return r.Msg
}
