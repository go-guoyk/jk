package jk

import (
	"context"
	"fmt"
	"net/url"
)

type Job struct {
	Class string `json:"_class"`
	Name  string `json:"name"`
	URL   string `json:"url"`
	Color string `json:"color"`
}

func (c *Client) JobGetConfig(ctx context.Context, name string) (cfg []byte, err error) {
	cfg, err = c.XMLGet(ctx, fmt.Sprintf("/job/%s/config.xml", name), nil)
	return
}

func (c *Client) JobCreateByConfig(ctx context.Context, name string, cfg []byte) (out []byte, err error) {
	q := url.Values{}
	q.Set("name", name)
	out, err = c.XMLPost(ctx, "/createItem", q, cfg)
	return
}
