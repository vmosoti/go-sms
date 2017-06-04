package twilio

import (
	"testing"
	"github.com/stretchr/testify/require"
)

const testClientID = "testClientID"

func TestProvider_Get(t *testing.T) {
	t.Run("It will return a request with the proper url", func(t *testing.T){
		provider := &Provider{ ClientID: testClientID }
		req, err := provider.SMS("","", "")
		require.Nil(t, err)

		const expectedURI = "https://api.twilio.com/2010-04-01/Accounts/testClientID/Messages.json"
		require.Equal(t, req.URL.String(), expectedURI)

	})
}

func TestProvider_SMS(t *testing.T) {
	t.Run("It will return a request with the proper url params encoded in the body", func(t *testing.T){
		provider := Provider{ ClientID: testClientID }

		testCases := []struct{
			from, to, message string
		}{
			{ from: "one", to: "two", message: "three" },
		}

		for _, test := range testCases {
			req, err := provider.SMS(test.from, test.to, test.message)
			require.Nil(t, err)

			// because this request is not actually being processed after passing
			// through our http.Client, it will not yet have these necessary preconfigured
			// headers, so inject them here for testing
			req.Header.Set("Accept", "application/json")
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			require.Nil(t, req.ParseForm())

			require.Equal(t, test.to, req.Form.Get("To"))
			require.Equal(t, test.from, req.Form.Get("From"))
			require.Equal(t, test.message, req.Form.Get("Body"))
		}
	})
}

func TestProvider_MMS(t *testing.T) {
	t.Run("It will return a request with the proper url params", func(t *testing.T){
		provider := Provider{}

		testCases := []struct{
			from, to, message, mediaURL string
		}{
			{ from: "four", to: "five", message: "six", mediaURL: "seven" },
		}

		for _, test := range testCases {
			req, err := provider.MMS(test.from, test.to, test.message, test.mediaURL)
			require.Nil(t, err)

			// because this request is not actually being processed after passing
			// through our http.Client, it will not yet have these necessary preconfigured
			// headers, so inject them here for testing
			req.Header.Set("Accept", "application/json")
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			require.Nil(t, req.ParseForm())

			require.Equal(t, test.to, req.Form.Get("To"))
			require.Equal(t, test.from, req.Form.Get("From"))
			require.Equal(t, test.message, req.Form.Get("Body"))
			require.Equal(t, req.Form.Get("MediaUrl"), test.mediaURL)
		}
	})
}