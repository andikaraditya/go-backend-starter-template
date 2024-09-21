package api

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

var (
	ErrPayload           = errors.New("something wrong with payload format")
	ErrQuery             = errors.New("something wrong with query format")
	ErrUnauthorized      = errors.New("please signin")
	ErrUnverified        = errors.New("please verify your account")
	ErrInsufficientRoles = errors.New("insufficient roles")
	ErrNotFound          = errors.New("not found")
	ErrNoAccess          = errors.New("has no access")
	ErrNotImplemented    = errors.New("not implemented")
	ErrFailedValidation  = errors.New("failed validation")
)

type Response struct {
	Results any   `json:"results"`
	Total   int64 `json:"total"`
}

func GetUserId(c *fiber.Ctx) string {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["user_id"].(string)
}

func GetUpdatedField(req []byte) ([]string, error) {
	var f []string

	m := make(map[string]any)

	if err := json.Unmarshal(req, &m); err != nil {
		return nil, err
	}

	for s := range m {
		f = append(f, s)
	}

	return f, nil
}

func SendError(ctx *fiber.Ctx, statusCode int, errMsg string, err error) error {
	log.Error().
		Str("method", ctx.Method()).
		Str("path", ctx.Path()).
		Err(err).
		Str("msg", errMsg).
		Msg("")

	return ctx.Status(statusCode).JSON(fiber.Map{
		"status_code": statusCode,
		"message":     errMsg,
	})
}
