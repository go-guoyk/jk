package jk

import "context"

type Hudson struct {
	Class           string `json:"_class"`
	Mode            string `json:"mode"`
	NodeDescription string `json:"nodeDescription"`
	NodeName        string `json:"nodeName"`
	NumExecutors    int    `json:"numExecutors"`
	Description     string `json:"description"`
	Jobs            []Job  `json:"jobs"`
	URL             string `json:"url"`
	Views           []View `json:"views"`
}

func (c *Client) HudsonGet(ctx context.Context) (h Hudson, err error) {
	err = c.JSONGet(ctx, "/api/json", nil, &h)
	return
}
