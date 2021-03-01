package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	echo "github.com/labstack/echo/v4"
	jwk "github.com/lestrrat-go/jwx/jwk"
)

type cachedGoogleOauth2Certs struct {
	downloadedAt time.Time
	keySet       jwk.Set
}

var publicKeysURL = "https://www.googleapis.com/oauth2/v3/certs"
var cachedCerts cachedGoogleOauth2Certs

// GetKey godoc
// @Summary This function returns the key used to verify the given token.
// @Param token header string true "Unverified token"
func GetKey(token *jwt.Token) (interface{}, error) {
	t := time.Now()
	elapsed := t.Sub(cachedCerts.downloadedAt)
	if cachedCerts.keySet == nil || int(elapsed.Minutes()) >= 60 {
		fmt.Println("Public keys are not available or older than 1 hour. Downloading them from " + publicKeysURL)

		keySet, err := jwk.Fetch(context.Background(), publicKeysURL)
		if err != nil {
			return nil, err
		}

		cachedCerts.keySet = keySet
		cachedCerts.downloadedAt = time.Now()
	}

	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("expecting JWT header to have string kid")
	}

	key, found := cachedCerts.keySet.LookupKeyID(keyID)

	if !found {
		return nil, fmt.Errorf("unable to find key %q", keyID)
	}

	var pubkey interface{}
	if err := key.Raw(&pubkey); err != nil {
		return nil, fmt.Errorf("Unable to get the public key. Error: %s", err.Error())
	}

	return pubkey, nil
}

// Login godoc
// @Summary Login user an ID token (JWT)
// @Description Login user using a ID token. The token is verified based on the rules defined on https://developers.google.com/identity/sign-in/web/backend-auth
// @Param id_token body string true "The ID token"
// @Success 200
// @Failure 401 {string} string "Unauthorized user"
// @Router /auth [post]
func Login(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)

	// do something with decoded claims
	claims := token.Claims.(jwt.MapClaims)
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}

	// The value of aud in the ID token must be equal to one of your app's client IDs.
	if claims["aud"] != os.Getenv("APPLICATION_CLIENT_ID") {
		return c.String(http.StatusUnauthorized, "Untrusted audience/client ID")
	}

	// The value of iss in the ID token must be equal to accounts.google.com or https://accounts.google.com.
	issuer := claims["iss"]
	if (issuer != "https://accounts.google.com") && (issuer != "accounts.google.com") {
		return c.String(http.StatusUnauthorized, "Untrusted issuer")
	}

	if claims["email_verified"] != true {
		return c.String(http.StatusUnauthorized, "Email is not verified")
	}

	return c.String(http.StatusOK, "Verified")
}
