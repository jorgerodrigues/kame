package tokens

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userId string) (string, error) {
	// Define the secret key used for signing the token
	secretKey := []byte("your-secret-key")
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
		fmt.Println("Failed to sign the token:", err)
		return "", nil
	}
	fmt.Println("Generated JWT token:", signedToken)
	return signedToken, nil
}
