package plivo

import (
	"testing"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"encoding/json"
)

const testClientID = "testClientID"

func TestProvider_Get(t *testing.T) {
	t.Run("It will return a request with the proper url", func(t *testing.T){
		provider := &Provider{ ClientID: testClientID }
		req, err := provider.SMS("","", "")
		require.Nil(t, err)

		const expectedURI = "https://api.plivo.com/v1/Account/testClientID/Message/"
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

			contents, err := ioutil.ReadAll(req.Body)
			req.Body.Close()
			require.Nil(t, err)

			response := &SMSRequest{}
			err =json.Unmarshal(contents, response)
			require.Nil(t, err)

			require.Equal(t, test.from, response.From)
			require.Equal(t, test.to, response.To)
			require.Equal(t, test.message, response.Body)
		}
	})
}

func TestProvider_MMS(t *testing.T) {
	t.Run("It will return an error because Plivo doesn't support MMS.", func(t *testing.T){
		provider := Provider{}
		_, err := provider.MMS("one", "two", "three", "four")
		require.Error(t, err)
	})
}