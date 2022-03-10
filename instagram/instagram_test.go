package instagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thaitanloi365/go-social-auth/utils"
)

func TestLoginInstagram(t *testing.T) {
	var token = "EAANL91EzxSkBAOOzVcquOhpwRAKuUxnuLUN1TvpTBCZAOSAVe2vGj3IVZCmgOQikAR5ZCXK7VrT8ZBm9ZAvrQ6aYxhWJpyj9dr5qZCRD4fL66z7OhyCzDqLb7fSlYB6ZASS5hx00qCNSUteKHOsFB3kaYvF0OHppsSSmUOrrdtI7Djx0jvQK17zcXRZBTPZA7MU6lIHTMGy6YoH1q6GpuLUbqXBhl2Sex4wMcWuSP4ZATyDaRU1HEZA8iLl"
	var instagramauth = New().WithAppID("650961669571775")
	result, err := instagramauth.Login(token)
	assert.NoError(t, err)

	utils.PrintJSON(result)
}
