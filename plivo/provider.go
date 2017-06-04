package plivo

import (
	"errors"
	"net/http"
	"github.com/b3ntly/go-authhttp"
	"net/url"
	"bytes"
	"fmt"
	"encoding/json"
)

type (
	// A Provider represents a SaaS service by which messages may be sent
	Provider struct {
		ClientID     string
		ClientSecret string
	}

	// JSON request for the Plivo API, note that src/dst/text are the identifiers not are NOT intuitive
	SMSRequest struct {
		From string `json:"src"`
		To string `json:"dst"`
		Body string `json:"text"`
	}
)

const (
	URITemplate = "https://api.plivo.com/v1/Account/%v/%v/"
	DefaultResource = "Message"
)

// Headers is a series of headers present in EVERY http request made by this provider
func Headers() map[string]string {
	return map[string]string {
		"Accept": "application/json",
	}
}

// Return the unique name of the provider
func (plivo Provider) Name() string { return "plivo" }

// Return an http.Client configured to use this provider
func (plivo Provider) Get() (*http.Client, error){
	if plivo.ClientID == "" || plivo.ClientSecret == "" {
		return nil, errors.New("Failed to provide valid clientId and/or clientSecret to provider: plivo")
	}

	transportOptions := []authhttp.TransportOption{
		authhttp.WithBasicAuth(plivo.ClientID, plivo.ClientSecret),
	}

	for key, value := range Headers() {
		transportOptions = append(transportOptions, authhttp.WithHeader(key, value))
	}

	client := authhttp.NewHTTPClient(transportOptions...)
	return client, nil
}

// Return an (unsent) http.Request object which will send an SMS message with the given provider
func (plivo Provider) SMS(from, to, message string) (*http.Request, error){
	endpoint, err := url.Parse(fmt.Sprintf(URITemplate, plivo.ClientID, DefaultResource))

	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(&SMSRequest{ From: from, To: to, Body: message })
	if err != nil {
		return nil, err
	}

	return buildRequest("POST", body, endpoint)
}

// Return an (unsent) http.Request object which will send an MMS message with the given provider
func (plivo Provider) MMS(from, to, message, mediaUrl string) (*http.Request, error){
	return nil, errors.New("Plivo does not support MMS messages.")
}

func buildRequest(verb string, body []byte, endpoint *url.URL) (*http.Request, error) {
	rb := *bytes.NewReader(body)
	req, err := http.NewRequest(verb, endpoint.String(), &rb)

	if err != nil {
		return nil, err
	}

	return req, nil
}