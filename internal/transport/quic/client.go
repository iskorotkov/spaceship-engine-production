package quic

import (
	"context"
	"crypto/tls"
	"encoding/gob"
	"errors"
	"io"
	"time"

	"github.com/iskorotkov/spaceship-engine-production/internal/transport"
	"github.com/lucas-clemente/quic-go"
)

var _ = (transport.Client)(Client{})

type Client struct {
	session quic.Session
	stream  quic.Stream
	encoder *gob.Encoder
	decoder *gob.Decoder
}

func NewClient(addr string, tlsConfig *tls.Config) (Client, error) {
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

	return Client{session, stream, encoder, decoder}, nil
}

func (c Client) Send(t transport.Type, value interface{}) error {
	log.Printf("sending %v", value)
	return c.encoder.Encode(transport.Request{Type: t, Message: value})
}

func (c Client) Recv() (interface{}, error) {
	var resp transport.Response
	for {
		if err := c.decoder.Decode(&resp); errors.Is(err, io.EOF) {
			log.Printf("waiting for response")

			time.Sleep(time.Second)
			continue
		} else if err != nil {
			return nil, err
		}

		log.Printf("got response %v", resp.Message)

		return resp.Message, nil
	}
}

func (c Client) Close() error {
	if err := c.session.CloseWithError(0, "client is closing"); err != nil {
		return err
	}

	return c.stream.Close()
}
