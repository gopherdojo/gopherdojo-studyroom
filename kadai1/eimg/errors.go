package eimg

// Default errors
var (
	ErrInvalidArgs = &Error{
		Name:        "invalid arguments",
		Description: "They are few arguments",
		Hint:        "arguments must be more than 1",
	}
	ErrInvalidPath = &Error{
		Name:        "invalid path",
		Description: "This path is invalid",
		Hint:        "Check if the path exists",
	}
)

// Error is a representation of errors returned from this package.
type Error struct {
	// Name is the name of this error.
	Name string `json:"error"`
	// Description is the description of this error.
	Description string `json:"decscription"`
	// Hint gives user further information.
	Hint string `json:"hint,omitempty"`
	// Debug gives debug information about this error.
	Debug string `json:"debug",omitempty`
}

// Error implement error interface
func (e *Error) Error() string {
	return e.Name
}

// WithHint updates hint
func (e *Error) WithHint(hint string) *Error {
	err := *e
	err.Hint = hint
	return &err
}

// WithDebug updates debug information
func (e *Error) WithDebug(debug string) *Error {
	err := *e
	err.Debug = debug
	return &err
}
