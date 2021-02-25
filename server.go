package main

import (
	"fmt"
	"github.com/foolin/goview/supports/echoview"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/stripe/stripe-go"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Renderer = echoview.Default()
	e.Static("/", os.Getenv("STATIC_DIR"))
	e.File("/", os.Getenv("STATIC_DIR")+"index.html")
	e.File("/manual", os.Getenv("STATIC_DIR")+"manual.html")
	e.File("/verify", os.Getenv("STATIC_DIR")+"verify.html")
	e.GET("/public-keys", publicKeyHandler)
	e.Logger.Fatal(e.Start("localhost:4242"))
}

type PublicKeys struct {
	stripeKey string `json:"stripe_key"`
	plaidkey string `json:"plaid_key"`
}

func publicKeyHandler( c echo.Context) (err error)  {

	data := PublicKeys{

	}
	return c.JSON(http.StatusOK, data)

}
