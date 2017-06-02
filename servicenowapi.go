//Provides an interface to interact with a ServiceNow server API
package servicenowapi

import (
	"errors"
	"log"
	"net/url"
	"fmt"
	"bytes"
	"encoding/base64"
)

// htmlMethods enum ------
type htmlMethods string

const (
	get    htmlMethods = "GET"
	put    htmlMethods = "PUT"
	post   htmlMethods = "POST"
	patch  htmlMethods = "PATCH"
	delete htmlMethods = "DELETE"
)

//------------------------

// htmlResponses enum -----
type htmlResponse uint

const (
	successBody      htmlResponse = 200
	created          htmlResponse = 201
	successNoBody    htmlResponse = 204
	badRequest       htmlResponse = 400
	unauthorized     htmlResponse = 401
	forbidden        htmlResponse = 403
	notFound         htmlResponse = 404
	methodNotAllowed htmlResponse = 405
)

//-------------------------

// Used to hold query parameters for requests
type QueryParameter struct {
	Key   string
	Value string
}

// Storage for a ServiceNow api connection
type apiConnection struct {
	baseURL     url.URL
	connected   bool
	sessionKey  string
	credentials string
}

func NewAPIConnection(URL string) *apiConnection {
	con := new(apiConnection)
	con.SetBaseURL(URL)
	con.connected = false

	return con
}

func (con *apiConnection) Connect(username string, password string) error {
	if con.baseURL == *new(url.URL) {
		return errors.New("Missing URL: Set a URL value before calling Connect()")
	}

	//request = http.NewRequest("GET")

	//resp, err := http.Get("%s://%s%s", con.baseURL.Scheme, con.baseURL.Host, "/manifest")

	encodedCreds := new(bytes.Buffer)
	base64.NewEncoder(base64.StdEncoding, encodedCreds).Write([]byte(username + ":" + password))
	con.credentials = encodedCreds.String()

	con.connected = true

	return nil
}

func (con *apiConnection) SetBaseURL(baseURL string) {
	u, err := url.ParseRequestURI(baseURL)
	if err != nil {
		log.Fatal(err)
	}
	con.baseURL = *u
	con.baseURL.Path = "/api/now/"
}

func (con *apiConnection) GetBaseURL() string {
	return fmt.Sprintf("%s://%s", con.baseURL.Scheme, con.baseURL.Host)
}

func (con *apiConnection) IsConnected() bool {
	return con.connected
}

func restRequest(con *apiConnection) error {
	return nil
}
