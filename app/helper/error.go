package helper

// ErrorPanic is a helper function to panic if an error is not nil.
func ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}
