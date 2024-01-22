package tokens

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

	var secretKey = []byte("your-secret-key")

func GenerateJWT(userId string, secret []byte) (string, error) {
	// Define the secret key used for signing the token
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)
	// Set the claims for the token
  
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "PAT" // Subject
  claims["iss"] = "upkame" // Issuer
  claims["user_id"] = userId // User ID
	// Sign the token with the secret key
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", nil
	}
	return signedToken, nil
}

func ValidateJWT(token string, secret []byte) (bool, error) {
  // Parse the token
  parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
    // Check the signing method
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
    }
    // Return the secret key used for signing
    return secret, nil
  })
  if err != nil {
    return false, nil
  }
  // Check if the token is valid
  if !parsedToken.Valid {
    return false, nil
  }
  return true, nil
}
