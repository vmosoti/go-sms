package sms

import (
	"net/http"
)

type (
	// A Provider represents a SaaS service by which messages may be sent
	Provider interface {
		// Returns the unique name of the provider
		Name() string

		// Return an http.Client configured for authenticated requests with the given provider
		Get() (*http.Client, error)

		// Provider specific request for sending a sms message
		SMS(from string, to string, message string) (*http.Request, error)

		// Provider specific request for sending an mms message
		MMS(from string, to string, message string, mediaUrl string) (*http.Request, error)
	}

	// Client interface the package sms
	Client struct {
		Provider Provider
		HTTPClient *http.Client
		number string
	}

	// Sender interface allows for simplified and intuitive message calls from a given phoneNumber
	Sender struct {
		Client *Client
		phoneNumber string
	}
)

// returns a Client configured for a given provider
func WithProvider(provider Provider) (*Client, error) {
	HTTPClient, err := provider.Get()

	if err != nil {
		return nil, err
	}

	return &Client{ Provider: provider, HTTPClient: HTTPClient }, nil
}

// returns a Sender configured for a given phoneNumber
func (c *Client) From(phoneNumber string) *Sender {
	return &Sender{ Client: c, phoneNumber: phoneNumber }
}

// returns a *http.Response of a type belonging to  the configured provider
func (s *Sender) SMS(to string, message string) (*http.Response, error) {
	req, err := s.Client.Provider.SMS(s.phoneNumber, to, message)

	if err != nil {
		return nil, err
	}

	return s.Client.HTTPClient.Do(req)
}

// returns a *http.Response of a type belonging to the configured provider
func (s *Sender) MMS(to string, message string, mediaUrl string) (*http.Response, error){
	req, err := s.Client.Provider.MMS(s.phoneNumber, to, message, mediaUrl)

	if err != nil {
		return nil, err
	}

	return s.Client.HTTPClient.Do(req)
}