package twilio

import (
	"github.com/b3ntly/go-authhttp"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"fmt"
)

// A Provider represents a SaaS service by which messages may be sent
type Provider struct {
	ClientID     string
	ClientSecret string
}

const (
	URITemplate = "https://api.twilio.com/2010-04-01/Accounts/%v/%v"
	DefaultResource = "Messages.json"
)

// Headers is a series of headers present in EVERY http request made by this provider
func Headers() map[string]string {
	return map[string]string {
		"Accept": "application/json",
		"Content-Type": "application/x-www-form-urlencoded",
	}
}

// Return the unique name of the provider
func (twilio Provider) Name() string { return "twilio" }

// Get returns an http.Client configured for the given provider
func (twilio Provider) Get() (*http.Client, error) {
	if twilio.ClientID == "" || twilio.ClientSecret == "" {
		return nil, errors.New("Failed to properly configure provider: Twilio.")
	}

	transportOptions := []authhttp.TransportOption{
		authhttp.WithBasicAuth(twilio.ClientID, twilio.ClientSecret),
	}

	for key, value := range Headers() {
		fmt.Println(key, value)
		transportOptions = append(transportOptions, authhttp.WithHeader(key, value))
	}

	client := authhttp.NewHTTPClient(transportOptions...)
	return client, nil
}

// SMS returns an *http.Request that will result in an SMS message being sent by the given provider
func (twilio Provider) SMS(from, to, message string) (*http.Request, error) {
	params := &url.Values{}
	params.Set("From", from)
	params.Set("To", to)
	params.Set("Body", message)
	endpoint, err := url.Parse(fmt.Sprintf(URITemplate, twilio.ClientID, DefaultResource))

	if err != nil {
		return nil, err
	}

	return buildRequest("POST", params, endpoint)
}

// SMS returns an *http.Request that will result in an SMS message being sent by the given provider
func (twilio Provider) MMS(from, to, message, mediaURL string) (*http.Request, error) {
	params := &url.Values{}
	params.Set("From", from)
	params.Set("To", to)
	params.Set("Body", message)
	params.Set("MediaUrl", mediaURL)

	endpoint, err := url.Parse(fmt.Sprintf(URITemplate, twilio.ClientID, DefaultResource))

	if err != nil {
		return nil, err
	}

	return buildRequest("POST", params, endpoint)
}

func buildRequest(verb string, params *url.Values, endpoint *url.URL) (*http.Request, error) {
	rb := *strings.NewReader(params.Encode())
	req, err := http.NewRequest(verb, endpoint.String(), &rb)

	if err != nil {
		return nil, err
	}

	return req, nil
}
