package instagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thaitanloi365/go-social-auth/utils"
)

func TestLoginInstagram(t *testing.T) {
	var token = "IGQVJWdXN0Q0NoTXpKczlqNDFZARzczd1pid1NPanVJM0JVRE5JQVZA5NEVSbkFqN2tPWTdJLXhBTG1YUW5HOUZAMRDZAMOWNWdTY3NGpJYVBlRkpzN2EyRnNrNGhLaDB3SHhhdC15TktjUGhwaDNBYXhHN0t5ZAURseGZAZAWnRv"
	var instagramauth = New()
	result, err := instagramauth.Login(token)
	assert.NoError(t, err)

	utils.PrintJSON(result)
}

func TestGetAccessTokenInstagram(t *testing.T) {
	var code = "AQClPRTwVutS9LC-bEC4y1Ms1X1i01ETGNLPFByYsRFqIBgfS5Ytwy6kzv2wAz_eRS3zsqmzz2IRPZsYFrRCRjityuFl5WdTOh5iGEUby7CeoPRAVF70ro-3y_gPmAifvACnBB_pqMrzTvx0cW8enApvT2KHXPV-BYvYzd_7KdsjRssKaOf99YUw-eJ2lGQAdSUMONUEKjv9NhLKtN6TvEr62bSGsN2wsyEuMi-AiuNOtQ"
	var redirectUri = "https://seller-staging.ezielog.com/"
	var instagramauth = New().WithAppID("test").WithAppSecret("test")
	result, err := instagramauth.GetAccessToken(code, redirectUri)
	assert.NoError(t, err)

	utils.PrintJSON(result)
}
