package nats

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"
	"time"

	"github.com/iskorotkov/spaceship-engine-production/internal/transport"
	"github.com/nats-io/nats.go"
)

var _ = (transport.Client)(&Client{})

type Client struct {
	nats     *nats.Conn
	lastResp interface{}
	m        sync.RWMutex
}

func NewClient(addr string) (Client, error) {
	nc, err := nats.Connect(addr)
	if err != nil {
		return Client{}, fmt.Errorf("error creating nats connection: %w", err)
	}

	log.Printf("nats connection created")

	return Client{
		nats:     nc,
		lastResp: nil,
		m:        sync.RWMutex{},
	}, nil
}

func (c *Client) Send(t transport.Type, value interface{}) error {
	var b bytes.Buffer
	encoder := gob.NewEncoder(&b)

	if err := encoder.Encode(transport.Request{Type: "", Message: value}); err != nil {
		return fmt.Errorf("error marshaling request to json: %w", err)
	}

	natsResp, err := c.nats.Request(t, b.Bytes(), time.Minute)
	if err != nil {
		return fmt.Errorf("error publishing request: %w", err)
	}

	log.Printf("nats client sent request")

	decoder := gob.NewDecoder(bytes.NewReader(natsResp.Data))

	var resp transport.Response
	if err := decoder.Decode(&resp); err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	c.m.Lock()
	defer c.m.Unlock()

	c.lastResp = resp.Message

	return nil
}

func (c *Client) Recv() (interface{}, error) {
	c.m.RLock()
	defer c.m.RUnlock()

	log.Printf("got nats response %v", c.lastResp)

	return c.lastResp, nil
}

func (c *Client) Close() error {
	c.nats.Close()

	log.Printf("nats connection closed")

	return nil
}
