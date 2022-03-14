package instagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thaitanloi365/go-social-auth/utils"
)

func TestLoginInstagram(t *testing.T) {
	var token = "AQBVsrUQLpARvTXGyXtvKndxYeDRLANlYBZWEQCZy1DwqiAhbS5FfY9Sg_ArdCvl58zGesQUKjK9Eip5JaExM5VnlGAldMkvfg1aAZkje0WCQcXdi8kS9YX9E3XqMs6-HO0uaD7wO2ak9i7rK6EthH8STsVBLbZ6xYSjRO_0clBHkMmRQqqaWmveCQUv0ZAbOQxf2ILqrJTY_CXry1bp-jL_4QKi1tLK3geRmlDOyk2L_w"
	var instagramauth = New().WithAppID("650961669571775")
	result, err := instagramauth.Login(token)
	assert.NoError(t, err)

	utils.PrintJSON(result)
}
