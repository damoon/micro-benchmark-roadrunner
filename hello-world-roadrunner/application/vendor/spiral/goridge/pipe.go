package goridge

import (
	"io"
	"sync"
	"errors"
)

// PipeRelay communicate with underlying process using standard streams (STDIN, STDOUT). Attention, use TCP alternative for
// Windows as more reliable option. This relay closes automatically with the process.
type PipeRelay struct {
	// How many bytes to write/read at once.
	BufferSize uint64

	mur sync.Mutex // concurrent read
	in  io.ReadCloser

	muw sync.Mutex // concurrent write
	out io.WriteCloser
}

// NewPipeRelay creates new pipe based data relay.
func NewPipeRelay(in io.ReadCloser, out io.WriteCloser) *PipeRelay {
	return &PipeRelay{BufferSize: BufferSize, in: in, out: out}
}

// Send signed (prefixed) data to underlying process.
func (rl *PipeRelay) Send(data []byte, flags byte) (err error) {
	rl.muw.Lock()
	defer rl.muw.Unlock()

	prefix := NewPrefix().WithFlags(flags).WithSize(uint64(len(data)))
	if _, err := rl.out.Write(prefix[:]); err != nil {
		return err
	}

	if _, err := rl.out.Write(data); err != nil {
		return err
	}

	return nil
}

// Receive data from the underlying process and returns associated prefix or error.
func (rl *PipeRelay) Receive() (data []byte, p Prefix, err error) {
	rl.mur.Lock()
	defer rl.mur.Unlock()

	defer func() {
		if rErr, ok := recover().(error); ok {
			err = rErr
		}
	}()

	if _, err := rl.in.Read(p[:]); err != nil {
		return nil, p, err
	}

	if !p.Valid() {
		return nil, p, errors.New("invalid prefix (checksum)")
	}

	if !p.HasPayload() {
		return nil, p, nil
	}

	data = make([]byte, 0, p.Size())
	leftBytes := p.Size()
	buffer := make([]byte, min(uint64(cap(data)), rl.BufferSize))

	for {
		if n, err := rl.in.Read(buffer); err == nil {
			data = append(data, buffer[:n]...)
			leftBytes -= uint64(n)
		} else {
			return nil, p, err
		}

		if leftBytes == 0 {
			break
		}
	}

	return
}

// Close the connection. Pipes are closed automatically with the underlying process.
func (rl *PipeRelay) Close() error {
	return nil
}
