package supersaas

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"sync"
)

const (
	defaultHost = "http://localhost:3000"
	apiVersion  = "1"
	pkgVersion  = "1.0.0"
)

type configuration struct {
	AccountName string
	Password    string
	Host        string
	DryRun      bool
	Verbose     bool
}

func defaultConfiguration() configuration {
	c := configuration{}
	c.AccountName = os.Getenv("SSS_API_ACCOUNT_NAME")
	c.Password = os.Getenv("SSS_API_PASSWORD")
	c.Host = defaultHost
	c.DryRun = false
	c.Verbose = false
	return c
}

// Client ...
type Client struct {
	configuration

	LastRequest *http.Request
	client      *http.Client

	Appointments *Appointments
	Forms        *Forms
	Schedules    *Schedules
	Users        *Users
}

var instance *Client
var once sync.Once

// GetInstance returns thread-safe singleton client instance ...
func GetInstance() *Client {
	once.Do(func() {
		instance = &Client{configuration: defaultConfiguration()}
	})
	return instance
}

// Configure sets auth credentials and runtime options for client
func (c *Client) Configure(accountName string, password string, host string, dryRun bool, verbose bool) {
	c.AccountName = accountName
	c.Password = password
	c.Host = host
	if c.Host == "" {
		c.Host = defaultHost
	}
	c.DryRun = dryRun
	c.Verbose = verbose
}

// NewClient creates a new Client with default configuration.
func NewClient(accountName string, password string) *Client {
	client := Client{configuration: defaultConfiguration(), client: http.DefaultClient}
	client.AccountName = accountName
	client.Password = password
	client.Appointments = &Appointments{API{Client: &client}}
	client.Forms = &Forms{API{Client: &client}}
	client.Schedules = &Schedules{API{Client: &client}}
	client.Users = &Users{API{Client: &client}}

	return &client
}

func userAgent() string {
	return "SSS/" + pkgVersion + " Go/" + runtime.Version() + " API/" + apiVersion
}

func (c Client) basicAuth() string {
	str := c.AccountName + ":" + c.Password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(str))
}

func (c Client) request(method string, path string, p map[string]interface{}, q map[string]interface{}, out interface{}) error {
	requestURL, err := url.Parse(c.Host + "/api" + path)
	if err != nil {
		return err
	}

	if len(q) > 0 {
		params := url.Values{}
		for k := range q {
			v := fmt.Sprint(q[k])
			params.Add(k, v)
		}
		requestURL.RawQuery = params.Encode()
	}

	var b io.ReadWriter
	if p != nil {
		b = new(bytes.Buffer)
		enc := json.NewEncoder(b)
		enc.SetEscapeHTML(false)
		err := enc.Encode(p)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, requestURL.String(), b)
	if err != nil {
		return err
	}
	c.LastRequest = req

	req.Header.Set("Accept", "application/json")
	if method != "GET" {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("User-Agent", userAgent())
	req.Header.Add("Authorization", c.basicAuth())

	if c.DryRun {
		return nil
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
	}

	if out != nil {
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(out)
		if err != nil {
			return err
		}
	}

	return nil
}

// Get performs an HTTP GET request
func (c Client) Get(path string, q map[string]interface{}, out interface{}) error {
	return c.request(http.MethodGet, path, nil, q, out)
}

// Post performs an HTTP POST request
func (c Client) Post(path string, p map[string]interface{}, q map[string]interface{}, out interface{}) error {
	return c.request(http.MethodPost, path, p, q, out)
}

// Put performs an HTTP PUT request
func (c Client) Put(path string, p map[string]interface{}, q map[string]interface{}, out interface{}) error {
	return c.request(http.MethodPut, path, p, q, out)
}

// Delete performs an HTTP DELETE request
func (c Client) Delete(path string, p map[string]interface{}, q map[string]interface{}) error {
	return c.request(http.MethodDelete, path, p, q, nil)
}
