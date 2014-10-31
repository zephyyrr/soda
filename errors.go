package soda

var (
	InvalidCode        = Error("Invalid code. Missing magic bytes.")
	UndefinedBehaviour = Error("undefined behviour detected.")
)

type Error string

func (e Error) Error() string {
	return string(e)
}
