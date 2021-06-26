package facebook

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thaitanloi365/go-social-auth/utils"
)

func TestLoginFacebook(t *testing.T) {
	var token = "EAAC2I0ITd0MBAARDwFhCZCY9WI7vTgkiQ5jjiRFyxKiO2vPeYvD2AzIR9TQoHoSZAkLWCttZCEtjV9bd4jWZBMvvFhkniPyYrG8HTSCIUPdcuo9JBpqVY2dOM4nskUK1vFeVPx5fTrtmoQOd3j0qZA98pzW1tT1584MUDuBkCtJQBBQhjtVgH3WZAde0BZCZANQjAbdp7WvIxhZCoV94rvUXz"
	var facebookauth = New().WithAppID("200262548682563")
	result, err := facebookauth.Login(token)
	assert.NoError(t, err)

	utils.PrintJSON(result)
}
