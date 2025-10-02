package rmq

import (
	"encoding/binary"
	"errors"

	"github.com/mikeziminio/rmq-envoy-filter/internal/ringbuffer"
)

const BufferLen = 140 * 1024 // 140 Kb

var ErrShortData = errors.New("short data")
var ErrInvalidFrameEnd = errors.New("invalid frame end")

type Logger interface {
	LogErrorf(format string, args ...interface{})
	LogInfof(format string, args ...interface{})
}

type RMQParser struct {
	rb          *ringbuffer.RingBuffer
	log         Logger
	insideFrame bool
}

func NewRMQParser(log Logger) *RMQParser {
	return &RMQParser{
		rb:  ringbuffer.NewRingBuffer(BufferLen),
		log: log,
	}
}

func (p *RMQParser) Parse(data []byte) ([]Frame, error) {
	_, err := p.rb.Write(data)
	if err != nil {
		return nil, err
	}

	if !p.insideFrame {
		frame, shift, err := p.parseFrame(data)
		if err != nil {
			if err == ErrShortData {

			} else if err == ErrInvalidFrameEnd {
				p.log.LogErrorf("invalid frame end")
			}
		}
	}

}

func (p *RMQParser) parseFrame(data []byte) (*Frame, int, error) {
	frameType := data[0]
	channel := binary.BigEndian.Uint16(data[1:3])
	payloadSize := binary.BigEndian.Uint32(data[3:7])

	totalSize := 7 + int(payloadSize) + 1
	if len(data) < totalSize {
		return nil, 0, ErrShortData
	}

	if data[7+payloadSize] != 0xCE {
		return nil, totalSize, ErrInvalidFrameEnd
	}

	frame := &Frame{
		Type:        frameType,
		Channel:     channel,
		PayloadSize: payloadSize,
		Payload:     data[7 : 7+payloadSize],
	}
	return frame, totalSize, nil
}
