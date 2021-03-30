package jk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type ClientOptions struct {
	URL        string
	Username   string
	Password   string
	HTTPClient *http.Client
}

type Client struct {
	opts ClientOptions
}

func New(opts ClientOptions) *Client {
	if strings.HasSuffix(opts.URL, "/") {
		opts.URL = strings.TrimSuffix(opts.URL, "/")
	}
	if opts.HTTPClient == nil {
		opts.HTTPClient = http.DefaultClient
	}
	return &Client{opts: opts}
}

func (c *Client) BuildURL(path string, query url.Values) string {
	u := c.opts.URL
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	u = u + path
	if query != nil {
		u = u + "?" + query.Encode()
	}
	return u
}

func (c *Client) JSONGet(ctx context.Context, path string, query url.Values, out interface{}) (err error) {
	var r *http.Request
	if r, err = http.NewRequestWithContext(ctx, http.MethodGet, c.BuildURL(path, query), nil); err != nil {
		return
	}
	r.Header.Set("Accept", "application/json")
	if c.opts.Username != "" && c.opts.Password != "" {
		r.SetBasicAuth(c.opts.Username, c.opts.Password)
	}
	var s *http.Response
	if s, err = c.opts.HTTPClient.Do(r); err != nil {
		return
	}
	defer s.Body.Close()
	if s.StatusCode >= 400 {
		var buf []byte
		buf, _ = ioutil.ReadAll(s.Body)
		err = fmt.Errorf("invalid status code: %d(%s)\n%s", s.StatusCode, s.Status, buf)
		return
	}
	err = json.NewDecoder(s.Body).Decode(out)
	return
}

func (c *Client) XMLGet(ctx context.Context, path string, query url.Values) (out []byte, err error) {
	var r *http.Request
	if r, err = http.NewRequestWithContext(ctx, http.MethodGet, c.BuildURL(path, query), nil); err != nil {
		return
	}
	r.Header.Set("Accept", "application/xml")
	if c.opts.Username != "" && c.opts.Password != "" {
		r.SetBasicAuth(c.opts.Username, c.opts.Password)
	}
	var s *http.Response
	if s, err = c.opts.HTTPClient.Do(r); err != nil {
		return
	}
	defer s.Body.Close()
	if out, err = ioutil.ReadAll(s.Body); err != nil {
		return
	}
	if s.StatusCode >= 400 {
		err = fmt.Errorf("invalid status code: %d(%s)\n%s", s.StatusCode, s.Status, out)
		return
	}
	return
}

func (c *Client) XMLPost(ctx context.Context, path string, query url.Values, in []byte) (out []byte, err error) {
	var r *http.Request
	if r, err = http.NewRequestWithContext(ctx, http.MethodPost, c.BuildURL(path, query), bytes.NewReader(in)); err != nil {
		return
	}
	r.Header.Set("Content-Type", "application/xml")
	r.Header.Set("Accept", "application/xml")
	if c.opts.Username != "" && c.opts.Password != "" {
		r.SetBasicAuth(c.opts.Username, c.opts.Password)
	}
	var s *http.Response
	if s, err = c.opts.HTTPClient.Do(r); err != nil {
		return
	}
	defer s.Body.Close()
	if out, err = ioutil.ReadAll(s.Body); err != nil {
		return
	}
	if s.StatusCode >= 400 {
		err = fmt.Errorf("invalid status code: %d(%s)\n%s", s.StatusCode, s.Status, out)
		return
	}
	return
}

func (c *Client) VoidPost(ctx context.Context, path string, query url.Values) (out []byte, err error) {
	var r *http.Request
	if r, err = http.NewRequestWithContext(ctx, http.MethodPost, c.BuildURL(path, query), bytes.NewReader([]byte{})); err != nil {
		return
	}
	r.Header.Set("Content-Type", "text/plain")
	if c.opts.Username != "" && c.opts.Password != "" {
		r.SetBasicAuth(c.opts.Username, c.opts.Password)
	}
	var s *http.Response
	if s, err = c.opts.HTTPClient.Do(r); err != nil {
		return
	}
	defer s.Body.Close()
	if out, err = ioutil.ReadAll(s.Body); err != nil {
		return
	}
	if s.StatusCode >= 400 {
		err = fmt.Errorf("invalid status code: %d(%s)\n%s", s.StatusCode, s.Status, out)
		return
	}
	return
}
