package facebookauth

import (
	"fmt"
	"testing"
)

func TestLoginFacebook(t *testing.T) {
	var token = "EAACejEGlyekBAKoV7PZAIo3qj8e7VcQZCihlW0E9GnjlioIcIt8vEnAb9KNJcpi8zSCZA0H42ZAa7YvGz3ctgqMj78fZARYCBiPikLZBZCsAiGakXA2iXUT8JPVIoGlI0ot2fWlgk6ZAxzbZB1j5qWNden5pwN27GyHwh5XSizffDxVAl1PZCZAVfquLE6MfGcs2iP0ZC5UbNumE7NM50obI0HkPQS90pt0LYSdbsKf2qtBIZBAZDZD"
	var facebookauth = New()
	result, err := facebookauth.Login(token)
	if err != nil {
		panic(err)
	}

	fmt.Println("result", result)
}
