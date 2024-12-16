package worker

import (
	"encoding/json"
	"errors"
	"product-worker/client"
)

type Job struct {
	Action string
	Market string
	Data   json.RawMessage
}

func (j *Job) Run(c *client.Client) (err error) {
	api, exists := c.Api(j.Market)

	if !exists {
		err = errors.New("Unsupported market\n")
		return
	}

	switch j.Action {
	case "stop":
		err = api.Stop(j.Data)
	default:
		err = errors.New("Unavailable action")
	}

	return
}
