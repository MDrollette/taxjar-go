package taxjar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

type Client struct {
	*http.Client
	baseUri string
	token   string
	Debug   bool

	Categories CategoryService
	Rates      RateService
	Taxes      TaxService
}

func NewClient(token string) *Client {
	c := &Client{Client: &http.Client{}, token: token, baseUri: "https://api.taxjar.com/v2", Debug: false}
	c.Setup()
	return c
}

func (c Client) Get(url string, queryParams interface{}) ([]byte, error) {
	req, _ := http.NewRequest("GET", c.baseUri+url, nil)
	req.Header.Add("Authorization", "Bearer "+c.token)
	req.Header.Add("Accept", "application/json")
	addQueryParams(req, queryParams)

	if c.Debug {
		fmt.Printf("%s %s\n", req.Method, req.URL)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error: %s", resp.StatusCode)
	}
	return data, err
}

func (c Client) Post(url string, params interface{}) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(buffer).Encode(params); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.baseUri+url, buffer)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	if c.Debug {
		fmt.Printf("%s %s %s\n", req.Method, req.URL, buffer)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		fmt.Printf("Response: %s\n", data)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error: %s", resp.StatusCode)
	}

	return data, err
}

func (c *Client) Setup() {
	c.Categories = CategoryService{Repository: CategoryApi{client: c}}
	c.Rates = RateService{Repository: RateApi{client: c}}
	c.Taxes = TaxService{Repository: TaxApi{client: c}}
}

func addQueryParams(req *http.Request, params interface{}) {
	v, _ := query.Values(params)
	req.URL.RawQuery = v.Encode()
}
