package jet

import (
	"encoding/json"
	"errors"
	"github.com/go-fires/fires/debug"
	"github.com/go-fires/fires/generator"
	"github.com/go-fires/fires/serializer"
)

type Client struct {
	service string

	transporter   Transporter
	serializer    serializer.Serializer
	dataFormatter DataFormatter
	pathGenerator PathGenerator
	idGenerator   generator.Generator
}

type Option func(*Client)

func New(service string, opts ...Option) *Client {
	c := &Client{
		service: service,
	}

	c.With(opts...)
	c.init()

	return c
}

func WithTransporter(transporter Transporter) func(*Client) {
	return func(c *Client) {
		c.transporter = transporter
	}
}

func WithSerializer(serializer serializer.Serializer) func(*Client) {
	return func(c *Client) {
		c.serializer = serializer
	}
}

func WithDataFormatter(dataFormatter DataFormatter) func(*Client) {
	return func(c *Client) {
		c.dataFormatter = dataFormatter
	}
}

func WithPathGenerator(pathGenerator PathGenerator) func(*Client) {
	return func(c *Client) {
		c.pathGenerator = pathGenerator
	}
}

func WithIDGenerator(idGenerator generator.Generator) func(*Client) {
	return func(c *Client) {
		c.idGenerator = idGenerator
	}
}

func (c *Client) init() {
	if c.dataFormatter == nil {
		c.dataFormatter = DefaultDataFormatter
	}

	if c.serializer == nil {
		c.serializer = serializer.Json
	}

	if c.pathGenerator == nil {
		c.pathGenerator = DefaultPathGenerator
	}

	if c.idGenerator == nil {
		c.idGenerator = generator.UUID
	}
}

func (c *Client) With(opts ...Option) *Client {
	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Client) Invoke(name string, request interface{}, response interface{}) error {
	req := Request{
		Path:   c.pathGenerator.Generate(c.service, name),
		Params: request,
		ID:     c.idGenerator.Generate(),
	}

	data, err := c.serializer.Serialize(
		c.dataFormatter.EncodeRequest(req),
	)
	if err != nil {
		return err
	}

	recv, err := c.transporter.Send(data)
	if err != nil {
		return err
	}

	resp, err := c.dataFormatter.DecodeResponse(recv, c.serializer.Unserialize)
	if err != nil {
		return err
	}

	if resp.Error != nil {
		return resp.Error
	}

	if resp.Result == nil {
		return errors.New("result is invalid")
	}

	debug.Dump(req, c.dataFormatter.EncodeRequest(req), recv, resp)

	return decodeResult(resp.Result, response)
}

func decodeResult(src interface{}, dest interface{}) error {
	result, err := json.Marshal(src)
	if err != nil {
		return err
	}

	return json.Unmarshal(result, dest)
}
