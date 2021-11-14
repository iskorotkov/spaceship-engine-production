package quic

import (
	"context"
	"crypto/tls"
	"encoding/gob"
	"errors"
	"io"
	"log"
	"time"

	"github.com/lucas-clemente/quic-go"
)

type Client struct {
	stream  quic.Stream
	encoder *gob.Encoder
	decoder *gob.Decoder
}

func New(addr string, tlsConfig *tls.Config) (Client, error) {
	tlsConfig.NextProtos = append(tlsConfig.NextProtos, "x-quic")

	session, err := quic.DialAddr(addr, tlsConfig, nil)
	if err != nil {
		return Client{}, err
	}

	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		return Client{}, err
	}

	encoder := gob.NewEncoder(stream)
	decoder := gob.NewDecoder(stream)

	return Client{stream, encoder, decoder}, nil
}

func (c Client) Send(t Type, value interface{}) error {
	log.Printf("sending %v", value)
	return c.encoder.Encode(Request{t, value})
}

func (c Client) Recv() (interface{}, error) {
	var resp Response
	for {
		if err := c.decoder.Decode(&resp); errors.Is(err, io.EOF) {
			log.Printf("waiting for response")

			time.Sleep(time.Second)
			continue
		} else if err != nil {
			return nil, err
		}

		log.Printf("got response")

		return resp.Message, nil
	}
}

func (c Client) Close() error {
	return c.stream.Close()
}
