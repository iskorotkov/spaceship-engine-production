package nats

import (
	"bytes"
	"crypto/tls"
	"encoding/gob"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/iskorotkov/spaceship-engine-production/internal/transport"
	"github.com/nats-io/nats.go"
)

var _ = (transport.Client)(&Client{})

type Client struct {
	nats     *nats.Conn
	lastResp *nats.Msg
	m        sync.RWMutex
}

func NewClient(addr string, config *tls.Config) (*Client, error) {
	nc, err := nats.Connect(addr, nats.Secure(config))
	if err != nil {
		return nil, fmt.Errorf("error creating nats connection: %w", err)
	}

	log.Printf("nats connection created")

	return &Client{
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

	resp, err := c.nats.Request(t, b.Bytes(), time.Minute)
	if err != nil {
		return fmt.Errorf("error publishing request: %w", err)
	}

	log.Printf("nats client sent request")

	c.m.Lock()
	defer c.m.Unlock()

	c.lastResp = resp

	return nil
}

func (c *Client) Recv() (interface{}, error) {
	c.m.RLock()
	defer c.m.RUnlock()

	if c.lastResp == nil {
		return nil, fmt.Errorf("no response available")
	}

	if err := c.lastResp.Header.Get("error"); err != "" {
		return nil, errors.New(err)
	}

	decoder := gob.NewDecoder(bytes.NewReader(c.lastResp.Data))

	var resp transport.Response
	if err := decoder.Decode(&resp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	log.Printf("got nats response %v", resp.Message)

	return resp.Message, nil
}

func (c *Client) Close() error {
	c.nats.Close()

	log.Printf("nats connection closed")

	return nil
}
