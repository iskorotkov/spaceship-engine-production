package tcp

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/iskorotkov/spaceship-engine-production/internal/transport"
)

var _ = (transport.Client)(Client{})

type Client struct {
	conn    net.Conn
	encoder *gob.Encoder
	decoder *gob.Decoder
}

func NewClient(addr string) (Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return Client{}, fmt.Errorf("error creating connection: %w", err)
	}

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	return Client{conn, encoder, decoder}, nil
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
	return c.conn.Close()
}
