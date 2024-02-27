package providers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReadNamesFromClaims(t *testing.T) {
	var testMatrix = []struct {
		rawData          map[string]interface{}
		ForeNameClaim    string
		SurNameClaim     string
		ForeNameExpected string
		SurNameExpected  string
	}{
		// read default claim
		{
			rawData: map[string]interface{}{
				DefaultForeNameClaim: "jhon",
				DefaultSurNameClaim:  "doe",
			},
			ForeNameClaim:    "",
			SurNameClaim:     "",
			ForeNameExpected: "jhon",
			SurNameExpected:  "doe",
		},
		// read customProvider claim
		{
			rawData: map[string]interface{}{
				"customProvider-forename-claim": "jhon",
				"customProvider-surname-claim":  "doe",
			},
			ForeNameClaim:    "customProvider-forename-claim",
			SurNameClaim:     "customProvider-surname-claim",
			ForeNameExpected: "jhon",
			SurNameExpected:  "doe",
		},
		// read customProvider claims that doesnt comes from idp...(bad mapping)
		{
			rawData:          map[string]interface{}{},
			ForeNameClaim:    "customProvider-forename-claim",
			SurNameClaim:     "customProvider-surname-claim",
			ForeNameExpected: "",
			SurNameExpected:  "",
		},
	}

	for _, ts := range testMatrix {
		forename, surname := ReadNamesFromClaims(ts.ForeNameClaim, ts.SurNameClaim, ts.rawData)
		assert.Equal(t, ts.ForeNameExpected, forename)
		assert.Equal(t, ts.SurNameExpected, surname)
	}
}

func Test_ReadEmailFromClaims(t *testing.T) {
	var testMatrix = []struct {
		rawData       map[string]interface{}
		emailClaim    string
		emailExpected string
	}{
		// read default claim
		{
			rawData: map[string]interface{}{
				DefaultEmailClaim: "jhon@doe.com",
			},
			emailExpected: "jhon@doe.com",
		},
		// read customProvider claim
		{
			rawData: map[string]interface{}{
				"customProvider-email-claim": "jhon@doe.com",
			},
			emailClaim:    "customProvider-email-claim",
			emailExpected: "jhon@doe.com",
		},
		// read customProvider claims that doesnt comes from idp...(bad mapping)
		{
			rawData:       map[string]interface{}{},
			emailClaim:    "customProvider-email-claim",
			emailExpected: "",
		},
		// WIF
		{
			rawData: map[string]interface{}{
				"http://schemas.xmlsoap.org/ws/2005/05/identity/claims/": "",
				WIFUniqueName: "jhon@doe.com",
			},
			emailClaim:    "",
			emailExpected: "jhon@doe.com",
		},
	}

	for _, ts := range testMatrix {
		email := ReadEmailFromClaims(ts.emailClaim, ts.rawData)
		assert.Equal(t, ts.emailExpected, email)
	}
}
