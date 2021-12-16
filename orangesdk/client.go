// Package orangesdk Orange Money Golang SDK.
//
// http://SERVER_IP:PORT/apigatewayom/apigwomService
//
// https://cdf.rdc.orange-money.com/tango/
//
// https://usd.rdc.orange-money.com/tango/
//
package orangesdk

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Client interface
type Client interface {
	Do(*http.Request) (*http.Response, error)
}

// ClientFunc function
type ClientFunc func(*http.Request) (*http.Response, error)

// Do Executes the http.Request
func (f ClientFunc) Do(r *http.Request) (*http.Response, error) {
	return f(r)
}

// Decorator is a function that takes a Client and returns a Client
type Decorator func(Client) Client

// Decorate function takes a client and some number of decorators and returns a
// decorated client
func Decorate(c Client, ds ...Decorator) Client {
	decorated := c
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}

// Header is a Client Decorator.
func Header(name, value string) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			r.Header.Add(name, value)
			return c.Do(r)
		})
	}
}

// APIClient struct
type APIClient struct {
	cli        Client
	token      string
	partnID    string
	mermsisdn  string
	partnName  string
	addr       string
	remoteIP   string
	remotePort string
	logger     *log.Logger
}

// address returns the remote address of the orange money service.
func (sdk *APIClient) address() string {
	return "https://" + sdk.remoteIP + ":" + sdk.remotePort + "/apigatewayom/apigwomService"
}

// Do to satisfy the http.Client interface.
func (sdk *APIClient) Do(r *http.Request) (*http.Response, error) {
	return sdk.cli.Do(r)
}

// Post function uses the client to post the data to the specified address
// and returns a byte slice and an error.
func (sdk *APIClient) Post(data io.Reader) ([]byte, error) {
	request, err := http.NewRequest(http.MethodPost, sdk.address(), data)
	if err != nil {
		sdk.logger.Fatalln(err)
	}
	resp, err := sdk.Do(request)
	if err != nil {
		sdk.logger.Fatalln(err)
	}
	resbytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		sdk.logger.Fatalln(err)
	}
	return resbytes, err
}

// NewClient returns a new APIClient
func NewClient(authToken, remoteIP, remotePort string) *APIClient {
	logger := log.New(os.Stdout, "drcorangeclient: ", log.Ldate|log.Ltime|log.Lshortfile|log.LstdFlags)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	return &APIClient{
		token:    authToken,
		remoteIP: remoteIP, remotePort: remotePort,
		logger: logger,
		cli: Decorate(
			client,
			Header("Accept", "application/xml,text/xml"),
			Header("Authorization", "Bearer "+authToken),
		),
	}
}
