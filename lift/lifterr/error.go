package lifterr

// LiftClientError type
type LiftClientError struct {
	embedded error
}

// NewLiftClientError func
func NewLiftClientError(err error) LiftClientError {
	return LiftClientError{
		embedded: err,
	}
}

// Error implementation
func (l LiftClientError) Error() string {
	return "lift client error: " + "\n\t" + l.embedded.Error()
}
