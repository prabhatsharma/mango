package auth

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
)

// KeySet can be used again and again
var KeySet = GetAWSCognitoJWK()

// GetAWSCognitoJWK will download the JWK from cognito website and store it so that it can be used later
func GetAWSCognitoJWK() jwk.Set {
	// docs at https://docs.aws.amazon.com/cognito/latest/developerguide/amazon-cognito-user-pools-using-tokens-verifying-a-jwt.html
	// https://cognito-idp.{region}.amazonaws.com/{userPoolId}/.well-known/jwks.json.
	// jwkURL := https://cognito-idp.us-west-2.amazonaws.com/us-west-2_VvrEQgmzT/.well-known/jwks.json

	// wellKnownURL := "https://cognito-idp.us-west-2.amazonaws.com/us-west-2_VvrEQgmzT/.well-known/jwks.json"
	wellKnownURL := os.Getenv("JWK_URL")
	if wellKnownURL == "" {
		return nil
	}

	fmt.Println("Reading JWK")

	fmt.Println("wellKnownURL: ", wellKnownURL)

	jwk.Fetch(context.Background(), wellKnownURL)

	keySet, errJwk := jwk.Fetch(context.Background(), wellKnownURL)
	if errJwk != nil {
		fmt.Println("Error fetching and parsing JWK: ", errJwk.Error())
		return nil
	}
	return keySet
}

func verifyToken(tokenString string) (*jwt.Token, error) {

	tkn, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != "RS256" { // jwa.RS256.String() works as well
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("kid header not found")
		}
		keys, _ := KeySet.LookupKeyID(kid)
		var raw interface{}
		return raw, keys.Raw(&raw)
	})

	if err != nil {
		fmt.Println("Error parsing: ", err.Error())
	}

	return tkn, err
}
