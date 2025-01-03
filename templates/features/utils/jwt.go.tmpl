package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

// GenerateToken generates a JWT token for a user with the provided user details.
// It includes user ID, email, first name, last name, and an expiration time of 72 hours.
//
// Parameters:
//   - userID: The unique identifier for the user.
//   - email: The email address of the user.
//   - firstName: The first name of the user.
//   - lastName: The last name of the user.
//
// Returns:
//   - A signed JWT token as a string.
//   - An error if there is an issue generating the token.
func GenerateToken(userID string, email string, firstName string, lastName string) (string, error) {
    claims := jwt.MapClaims{}
    claims["user_id"] = userID
    claims["email"] = email
    claims["first_name"] = firstName
    claims["last_name"] = lastName
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedToken, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }
    return signedToken, nil
}

// VerifyToken verifies a JWT token string and returns the claims if the token is valid.
// It parses the token, checks the signing method, and validates the token.
//
// Parameters:
//   - tokenString: A string representing the JWT token to be verified.
//
// Returns:
//   - jwt.MapClaims: A map containing the claims if the token is valid.
//   - error: An error if the token is invalid or if there is an issue during parsing or validation.
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
    // Parse the token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Check the signing method
        if token.Method != jwt.SigningMethodHS256 {
            return nil, fmt.Errorf("invalid signing method")
        }

        return secretKey, nil
    })

    // Check for errors
    if err != nil {
        return nil, err
    }

    // Validate the token
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, fmt.Errorf("invalid token")
}
