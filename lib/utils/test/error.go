package test

// ErrStr returns a string representation of the error, or empty if the error doesn't exist
func ErrStr(err error) string {
	if err == nil {
		return ""
	}

	return err.Error()
}
