package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var headers []map[string]string

type Client struct {
	baseUrl string
}

func NewHttpClient(baseUrl string) *Client {
	return &Client{
		baseUrl: baseUrl,
	}
}

func (c *Client) SetHeader(header map[string]string) {
	headers = append(headers, header)
}

func (c *Client) Get(path string) ([]byte, error) {
	data, err := c.build(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	dataByte, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return dataByte, nil
}

func (c *Client) Post(path string, payload interface{}) ([]byte, error) {
	dataByte, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	data, err := c.build(http.MethodPost, path, dataByte)
	fmt.Println("")
	fmt.Printf("HTTPCLIENT RES DATA => %v", data)
	fmt.Println("")
	if err != nil {
		return nil, err
	}
	respData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return respData, nil
}

func (c *Client) Put(path string, payload interface{}) ([]byte, error) {
	dataByte, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	data, err := c.build(http.MethodPut, path, dataByte)
	fmt.Printf("DATA => %v", data)
	if err != nil {
		return nil, err
	}
	respData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	return respData, nil
}

func (c *Client) Delete(path string) ([]byte, error) {
	data, err := c.build(http.MethodDelete, path, nil)

	if err != nil {
		return nil, err
	}
	respData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	return respData, nil
}

func (c *Client) build(method string, path string, payload []byte) (interface{}, error) {
	client := http.Client{}

	req, err := http.NewRequest(method, fmt.Sprintf("%v/%v", c.baseUrl, path), bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	for _, header := range headers {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// if resp.StatusCode > 300 {
	// 	return nil, errors.New("not found")
	// }

	var data interface{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
