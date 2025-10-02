package rmq

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
