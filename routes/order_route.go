package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtutils "github.com/hieronimusbudi/go-bookstore-utils/jwt"
	resterrors "github.com/hieronimusbudi/go-bookstore-utils/rest_errors"
	"github.com/hieronimusbudi/go-fiber-bookstore-order-api/controllers"
)

var (
	// jwtSecret     = os.Getenv("JWT_SECRET")
	// jwtCookieName = os.Getenv("JWT_COOKIE_NAME")
	jwtSecret     = "secret"
	jwtCookieName = "token::jwt"
)

func ValidateRequest(c *fiber.Ctx) error {
	// Get token from cookie
	token := c.Cookies(jwtCookieName)
	if token == "" {
		restJwtErr := resterrors.NewUnauthorizedError("Unauthorized")
		return c.Status(restJwtErr.Status()).JSON(restJwtErr)
	}

	// Validate token
	tokenClaims, tokenErr := jwtutils.ValidateToken(token, jwtSecret)
	if tokenErr != nil {
		restJwtErr := resterrors.NewUnauthorizedError("Token claims not exists")
		return c.Status(restJwtErr.Status()).JSON(restJwtErr)
	}

	c.Context().SetUserValue("tokenClaims", tokenClaims)
	return c.Next()
}

func OrderRoutes(app *fiber.App) {
	secureRoute := app.Group("/", ValidateRequest)
	secureRoute.Post("/api/orders", controllers.CreateOrder)
}
