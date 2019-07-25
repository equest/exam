package auth

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

const jwksURI = "https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json"

// JWK ...
type JWK struct {
	Keys []struct {
		Alg string `json:"alg"`
		E   string `json:"e"`
		Kid string `json:"kid"`
		Kty string `json:"kty"`
		N   string `json:"n"`
	} `json:"keys"`
}

// AWSCognitoAuthService ...
type AWSCognitoAuthService struct {
	JWK        *JWK
	URI        string
	Region     string
	UserPoolID string
}

// NewAWSCognitoAuthService creates new AWSCognitoAuthService
func NewAWSCognitoAuthService(region, poolID string) *AWSCognitoAuthService {
	s := &AWSCognitoAuthService{
		URI:        fmt.Sprintf(jwksURI, region, poolID),
		Region:     region,
		UserPoolID: poolID,
	}
	return s
}

// LoadJWK loads JWKS
func (s *AWSCognitoAuthService) LoadJWK() error {
	r, err := http.NewRequest("GET", fmt.Sprintf(jwksURI, s.Region, s.UserPoolID), nil)
	if err != nil {
		return err
	}
	r.Header.Add("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return nil
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	jwk := new(JWK)
	err = json.Unmarshal(body, jwk)
	if err != nil {
		return err
	}
	s.JWK = jwk
	return nil
}

//ParseJWT parses jwt token
func (s *AWSCognitoAuthService) ParseJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		kid := token.Header["kid"].(string)
		for _, key := range s.JWK.Keys {
			if key.Kid == kid {
				return convertKey(key.E, key.N), nil
			}
		}
		return nil, errors.New("Key ID not found")
	})
}

func convertKey(rawE, rawN string) *rsa.PublicKey {
	decodedE, err := base64.RawURLEncoding.DecodeString(rawE)
	if err != nil {
		panic(err)
	}
	if len(decodedE) < 4 {
		ndata := make([]byte, 4)
		copy(ndata[4-len(decodedE):], decodedE)
		decodedE = ndata
	}
	pubKey := &rsa.PublicKey{
		N: &big.Int{},
		E: int(binary.BigEndian.Uint32(decodedE[:])),
	}
	decodedN, err := base64.RawURLEncoding.DecodeString(rawN)
	if err != nil {
		panic(err)
	}
	pubKey.N.SetBytes(decodedN)
	return pubKey
}

// ValidateToken validate authentication token
func (s *AWSCognitoAuthService) ValidateToken(token string) error {
	t, err := s.ParseJWT(token)
	if err != nil {
		return err
	}
	if !t.Valid {
		return errors.New("unauthorized")
	}
	c := t.Claims.(jwt.MapClaims)
	tuse := c["token_use"]
	stuse := tuse.(string)
	if stuse != "access" {
		return errors.New("unauthroized")
	}
	return nil
}
