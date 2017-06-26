package route

// DispatchNonBlocking will dispatch in to buf.
// if buf is full, will discard the data
func dispatchNonBlocking(buf chan []byte, in []byte) {
	select {
	case buf <- in:
	default:
	}
}

// DispatchBlocking will dispatch in to buf.
// If buf is full, the call will block
func dispatchBlocking(buf chan []byte, in []byte) {
	buf <- in
}
