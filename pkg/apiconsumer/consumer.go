package apiconsumer

import (
	"io"
	"net/http"
)

type Consumer struct {
	client *http.Client
}

func NewConsumer() *Consumer {
	return &Consumer{
		client: &http.Client{},
	}
}

func (c *Consumer) GetStats(url string) (int, []byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return 0, nil, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, respBody, nil
}
