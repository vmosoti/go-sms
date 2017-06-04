## SMS

This project is a multi-provider client library for SMS and MMS messages written in Golang.

## Usage

```go 
package main

import (
    "github.com/b3ntly/sms"
    "github.com/b3ntly/sms/twilio"
)

func main(){
    provider, err := sms.WithProvider(twilio{ ClientID: "xxxYourClientID", ClientSecret: "xxxYourClientSecret" })
    client := client.From("xxxYourTwilioNumber")
    resp, err := client.SMS("xxxRecipientNumber", "hello")
}
```