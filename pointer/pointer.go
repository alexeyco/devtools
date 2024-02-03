package pointer

// Ptr returns pointer by value.
func Ptr[T any](v T) *T {
	return &v
}

// Val returns value by pointer.
func Val[T any](p *T) (v T) {
	if p != nil {
		v = *p
	}

	return
}
