package must

// Get returns the value v if the error is not nil,
// otherwise it panics.
func Get[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
