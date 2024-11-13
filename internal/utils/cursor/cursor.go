package cursor

import (
	"real-time-streaming/internal/utils/errList"
)

type Cursor interface {
	Consumer
	Producer
}

type Consumer interface {
	Next() bool                                   // Move to the next item
	GetCurrent() (interface{}, errList.CoreError) // Retrieve current data
	Close() errList.CoreError                     // Mark consumer as done
}

type Producer interface {
	SendData(record interface{}) bool // Send record or error to cursor
	IsClosed() bool                   // Check if cursor is closed
	Done()                            // Mark producer as done
}

type cursor struct {
	dataCh   chan data
	current  data
	isClosed bool
}

type data struct {
	record interface{}
	err    errList.CoreError
}

const defaultBufferSize = 10000

// NewCursor creates a cursor with default buffer size
func NewCursor() Cursor {
	return newCursor(defaultBufferSize)
}

// NewCursorWithBuffer allows custom buffer size
func NewCursorWithBuffer(bufferSize int) Cursor {
	if bufferSize <= 0 {
		bufferSize = defaultBufferSize
	}
	return newCursor(bufferSize)
}

func newCursor(bufferSize int) Cursor {
	return &cursor{
		dataCh: make(chan data, bufferSize),
	}
}

func (c *cursor) SendData(record interface{}) bool {
	if c.isClosed {
		return true
	}
	select {
	case c.dataCh <- data{record: record}:
	case <-c.isClosedCh():
	}
	return c.isClosed
}

func (c *cursor) Next() bool {
	item, ok := <-c.dataCh
	if ok {
		c.current = item
		return true
	}
	return false
}

func (c *cursor) GetCurrent() (interface{}, errList.CoreError) {
	return c.current.record, c.current.err
}

func (c *cursor) Close() errList.CoreError {
	if !c.isClosed {
		c.isClosed = true
		close(c.dataCh)
	}
	return errList.CoreError{}
}

func (c *cursor) Done() {
	c.Close()
}

func (c *cursor) IsClosed() bool {
	return c.isClosed
}

func (c *cursor) isClosedCh() chan struct{} {
	ch := make(chan struct{})
	go func() {
		if c.isClosed {
			close(ch)
		}
	}()
	return ch
}
