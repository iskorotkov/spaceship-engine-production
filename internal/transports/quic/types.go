package quic

type Type = string

type Request struct {
	Type    Type
	Message interface{}
}

type Response struct {
	Message interface{}
}
