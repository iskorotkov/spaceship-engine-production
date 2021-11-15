package transport

type Type = string

type Handler func(req interface{}) (Response, error)

type Request struct {
	Type    Type
	Message interface{}
}

type Response struct {
	Message interface{}
}

type Client interface {
	Send(t Type, value interface{}) error
	Recv() (interface{}, error)
	Close() error
}

type Server interface {
	Start(addr string) error
	Handle(t Type, h Handler) error
	Close() error
}
