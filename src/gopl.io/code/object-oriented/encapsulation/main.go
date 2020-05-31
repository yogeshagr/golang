/*

As one of the benefits of encapsulation, consider the bytes.Buffer type. It is
frequently used to accumulate very short strings, so it is a profitable
optimization to reserve a little extra space in the object to avoid memory
allocation in thsi common case. Since Buffer is a struct type, this space takes
the form of an extra field of type [64]byte with an uncapitalized name. when
this field was added, because it was not exported, clients of Buffer outside
the bytes package were unaware of any change except improved performance. Buffer
and its Grow method are show below, simplified for clarity.

*/

package main

type Buffer struct {
	buf     []byte
	initial [64]byte
	/* ... */
}

// Grow exapnds the buffer's capacity, if necessary,
// to guarantee space for another n bytes. [...]
func (b *Buffer) Grow(n int) {
	if b.buf == nil {
		b.buf = b.initial[:0] // use preallocated space initially
	}
	if len(b.buf)+n > cap(b.buf) {
		buf := make([]byte, len(b.buf), 2*cap(b.buf)+n)
		copy(buf, b.buf)
		b.buf = buf
	}
}
