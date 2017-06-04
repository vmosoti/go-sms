[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/b3ntly/sms/master/LICENSE.txt) 
[![Build Status](https://travis-ci.org/b3ntly/sms.svg?branch=master)](https://travis-ci.org/b3ntly/sms)
[![Coverage Status](https://coveralls.io/repos/github/b3ntly/sms/badge.svg?branch=master)](https://coveralls.io/github/b3ntly/sms?branch=master?q=1) 
[![GoDoc](https://godoc.org/github.com/b3ntly/sms?status.svg)](https://godoc.org/github.com/b3ntly/sms)

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