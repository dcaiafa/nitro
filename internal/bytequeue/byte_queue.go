package bytequeue

// ByteQueue is a queue of bytes.
// Data is added to the end of the queue via Write() or ReadNFrom(), and is
// removed via Pop(). The contents of the queue can be accessed via Peek().
type ByteQueue struct {
	//           offset
	//             v
	// +-----------+--------+-----------+
	// | available |  data  | available |
	// +-----------+--------+-----------+
	//             ^- used -^
	buf    []byte
	offset int
	used   int
}

// Write adds data to the end of the queue.
func (bq *ByteQueue) Write(data []byte) (int, error) {
	bq.ensure(len(data))

	copy(bq.buf[bq.offset+bq.used:], data[:])
	bq.used += len(data)

	return len(data), nil
}

// Peek returns a slice with the data available in the queue.
func (bq *ByteQueue) Peek() []byte {
	return bq.buf[bq.offset : bq.offset+bq.used]
}

// Pop removes n bytes from the beginning of the queue.
func (bq *ByteQueue) Pop(n int) {
	bq.offset += n
	bq.used -= n

	// If we've used all the data, we can safely move the offset back to the
	// beginning of the buffer.
	if bq.used == 0 {
		bq.offset = 0
	}
}

func (bq *ByteQueue) ensure(l int) {
	needed := bq.used + l
	if needed > len(bq.buf) {
		newSize := 2 * len(bq.buf)
		if newSize < needed {
			newSize = needed + (needed / 2)
		}

		newBuf := make([]byte, newSize)
		copy(newBuf[:], bq.Peek())
		bq.buf = newBuf
		bq.offset = 0
	} else if (bq.offset + bq.used + l) > len(bq.buf) {
		// Buffer is big enough, but the new data will not fit at the end. We must
		// move the existing data to the beginning of the buffer.
		copy(bq.buf[:], bq.buf[bq.offset:bq.offset+bq.used])
		bq.offset = 0
	}
}
