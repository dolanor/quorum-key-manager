package jsonrpc

import (
	"fmt"
	"reflect"
	"sync/atomic"

	httpclient "github.com/ConsenSysQuorum/quorum-key-manager/pkg/http/client"
)

var defaultVersion = "2.0"

type ClientConfig struct {
	Version string             `json:"version,omitempty"`
	HTTP    *httpclient.Config `json:"http,omitempty"`
}

func (cfg *ClientConfig) SetDefault() *ClientConfig {
	if cfg.HTTP == nil {
		cfg.HTTP = new(httpclient.Config)
	}

	cfg.HTTP.SetDefault()

	if cfg.Version == "" {
		cfg.Version = defaultVersion
	}

	return cfg
}

// Client is an jsonrpc client interface
type Client interface {
	// Do sends an jsonrpc request and returns an jsonrpc response, following
	// policy (such as redirects, cookies, auth)
	Do(*Request) (*Response, error)

	// CloseIdleConnections closes any connections on its Transport which
	// were previously connected from previous requests but are now
	// sitting idle in a "keep-alive" state. It does not interrupt any
	// connections currently in use.
	CloseIdleConnections()
}

// client is a connector to a jsonrpc server
type client struct {
	client httpclient.Client
}

// NewClient creates a new jsonrpc client from an HTTP client
func NewClient(c httpclient.Client) Client {
	return &client{
		client: c,
	}
}

// Do sends an jsonrpc request and returns an jsonrpc response
func (c *client) Do(req *Request) (*Response, error) {
	// write request body
	err := req.WriteBody()
	if err != nil {
		return nil, err
	}

	httpResp, err := c.client.Do(req.Request())
	if err != nil {
		return nil, err
	}

	// Create response and reads body
	resp := NewResponse(httpResp)
	err = resp.ReadBody()
	if err != nil {
		return resp, err
	}

	if req.ID() != nil {
		var respIDVal = reflect.New(reflect.TypeOf(req.ID()))
		err = resp.UnmarshalID(respIDVal.Interface())
		if err != nil {
			return resp, err
		}

		if respIDVal.Elem().Interface() != req.ID() {
			return resp, fmt.Errorf("request and response id didn't match")
		}
	}

	err = resp.Error()
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *client) CloseIdleConnections() {
	c.client.CloseIdleConnections()
}

type incrementalIDlient struct {
	client Client

	baseID    string
	idCounter uint32
}

// WithIncrementalID wraps a client with an ID counter an increases it each time a new request comes out
func WithIncrementalID(id interface{}) func(Client) Client {
	return func(c Client) Client {
		idC := &incrementalIDlient{
			client: c,
		}

		if id != nil {
			idC.baseID = fmt.Sprintf("%v.", id)
		}

		return idC
	}
}

func (c *incrementalIDlient) nextID() string {
	return fmt.Sprintf("%v%v", c.baseID, atomic.AddUint32(&c.idCounter, 1))
}

func (c *incrementalIDlient) Do(req *Request) (*Response, error) {
	req.WithID(c.nextID())
	return c.client.Do(req)
}

func (c *incrementalIDlient) CloseIdleConnections() {
	c.client.CloseIdleConnections()
}

type versionClient struct {
	client Client

	version string
}

// WithVersion wraps a client to set version each time a new request comes out
func WithVersion(version string) func(Client) Client {
	return func(c Client) Client {
		if version == "" {
			version = defaultVersion
		}
		return &versionClient{
			client:  c,
			version: version,
		}
	}
}

func (c *versionClient) Do(req *Request) (*Response, error) {
	req.WithVersion(c.version)
	return c.client.Do(req)
}

func (c *versionClient) CloseIdleConnections() {
	c.client.CloseIdleConnections()
}
