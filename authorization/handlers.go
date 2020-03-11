/*
Authors: JV Afable, JP Chanchico, Lance Lim

This is a course requirement for CS 192 Software Engineering II
under the supervision of Asst. Prof. Ma. Rowena C. Solamo of
the Department of Computer Science, College of Engineering,
University of the Philippines, Diliman for the AY 2019-2020.

Code History:
	March 10, 2020: Lance Lim - Initialized file.
*/

package auth

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

// jwtCustomClaims are custom claims extending default ones.
type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

// Login is the function that allows the admin to log in
func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	h := sha256.New()
	h.Write([]byte(password))

	hashedPassword := fmt.Sprintf("%x", h.Sum(nil))

	_ = godotenv.Load()

	adminUsername := os.Getenv("admin_username")
	adminPassword := os.Getenv("admin_password")
	jwtSecret := os.Getenv("jwt_secret")

	// Throws unauthorized error
	if username != adminUsername || hashedPassword != adminPassword {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		"Ad Min",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
