package jk

import (
	"context"
	"fmt"
	"net/url"
)

type View struct {
	Class       string `json:"_class"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Jobs        []Job  `json:"jobs"`
}

func (c *Client) ViewGet(ctx context.Context, name string) (view View, err error) {
	err = c.JSONGet(ctx, fmt.Sprintf("/view/%s/api/json", name), nil, &view)
	return
}

func (c *Client) ViewAddJob(ctx context.Context, name string, jobName string) (err error) {
	q := url.Values{}
	q.Set("name", jobName)
	_, err = c.VoidPost(ctx, fmt.Sprintf("/view/%s/addJobToView", name), q)
	return
}

func (c *Client) ViewRemoveJob(ctx context.Context, name string, jobName string) (err error) {
	q := url.Values{}
	q.Set("name", jobName)
	_, err = c.VoidPost(ctx, fmt.Sprintf("/view/%s/removeJobFromView", name), q)
	return
}
