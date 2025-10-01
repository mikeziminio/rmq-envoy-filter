package rmq

import (
	"encoding/binary"

	"github.com/proxy-wasm/proxy-wasm-go-sdk/proxywasm"
)

type FrameType byte

const (
	MethodFrame    FrameType = 1 // команда/метод (например Basic.Publish)
	HeaderFrame    FrameType = 2 //	Header frame — метаданные сообщения (body size, свойства)
	BodyFrame      FrameType = 3 //	Body frame — кусок тела сообщения
	HeartbeatFrame FrameType = 8 //	Heartbeat frame — пустой, для проверки живости соединения
)

type Frame struct {
	Type        byte
	Channel     uint16
	PayloadSize uint32
	Payload     []byte
}

func parseFrame(data []byte) (*Frame, int) {
	frameType := data[0]
	channel := binary.BigEndian.Uint16(data[1:3])
	payloadSize := binary.BigEndian.Uint32(data[3:7])

	totalSize := 7 + int(payloadSize) + 1
	if len(data) < totalSize {
		return nil, 0 // ждем остаток фрейма
	}

	if data[7+payloadSize] != 0xCE {
		proxywasm.LogErrorf("invalid frame end byte")
		return nil, 0
	}

	frame := &Frame{
		Type:        frameType,
		Channel:     channel,
		PayloadSize: payloadSize,
		Payload:     data[7 : 7+payloadSize],
	}
	return frame, totalSize
}
