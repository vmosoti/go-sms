package sms_test

import (
	"os"
	"testing"
	"github.com/b3ntly/sms"
	"github.com/b3ntly/sms/plivo"
	"github.com/b3ntly/sms/twilio"
	"fmt"
	"github.com/stretchr/testify/require"
)

type providerMetaData struct {
	clientID, clientSecret, testNumber, testRecipient string
}

const (
	attCustomerSupport = "+15417354469"
	dogPhoto = "https://cdn.pixabay.com/photo/2015/04/17/09/36/domestic-cat-726989_1280.jpg"
)

var metaData = map[string]*providerMetaData {
	"twilio": &providerMetaData{
		clientID: os.Getenv("twilioClientID"),
		clientSecret: os.Getenv("twilioClientSecret"),
		testNumber: os.Getenv("twilioTestNumber"),
		testRecipient: os.Getenv("attCustomerSupport"),
	},
	"plivo": &providerMetaData{
		clientID: os.Getenv("plivoClientID"),
		clientSecret: os.Getenv("plivoClientSecret"),
		testNumber: os.Getenv("plivoTestNumber"),
		testRecipient: os.Getenv("attCustomerSupport"),
	},
}

func TestClient(t *testing.T) {
	plivoProvider := plivo.Provider{ ClientID: metaData["plivo"].clientID, ClientSecret: metaData["plivo"].clientSecret  }
	twilioProvider := twilio.Provider{ ClientID: metaData["twilio"].clientID, ClientSecret: metaData["twilio"].clientSecret}

	providers := []sms.Provider {
		plivoProvider,
		twilioProvider,
	}

	for _, provider := range providers {
		if provider.Name() == "plivo" {
			continue // because their verification process isn't automated
		}

		t.Run(fmt.Sprintf("Provider: %v can send an SMS successfully", provider.Name()), func(t *testing.T){
			client, err := sms.WithProvider(provider)
			require.Nil(t, err)
			resp, err := client.From(metaData[provider.Name()].testNumber).SMS(attCustomerSupport, "hello")
			require.Nil(t, err)
			require.Equal(t, 201, resp.StatusCode)
		})

		t.Run(fmt.Sprintf("Provider: %v can send an MMS successfully unless it is plivo which doesn't support MMS", provider.Name()), func(t *testing.T){
			client, err := sms.WithProvider(provider)
			require.Nil(t, err)
			resp, err := client.From(metaData[provider.Name()].testNumber).MMS(attCustomerSupport, "hello", dogPhoto)

			// handle the edge case of Plivo not supporting MMS
			if provider.Name() == "plivo" {
				require.Error(t, err)
				return
			}

			require.Nil(t, err)
			defer resp.Body.Close()
			require.Equal(t, 201, resp.StatusCode)
		})
	}
}
