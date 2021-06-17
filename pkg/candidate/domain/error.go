package domain

const (
	CandidateDoesNotExists = int(0)
	EmailIsInvalid         = int(1)
)

type Error struct {
	Err  error
	Msg  string
	Type int
}

func (e Error) Unwrap() error {
	return e.Err
}

func (e Error) Error() string {
	return e.Msg
}

var ErrorCandidateNotExist = &Error{
	Msg:  "candidate doesn't exists",
	Type: CandidateDoesNotExists,
}

var InvalidEmail = &Error{
	Msg:  "email is invalid",
	Type: EmailIsInvalid,
}
