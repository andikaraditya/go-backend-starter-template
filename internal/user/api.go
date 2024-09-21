package user

import (
	"errors"

	"github.com/andikaraditya/go-backend-starter-template/internal/api"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {
	req := new(User)

	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	err := Service.CreateUser(*req)
	if err != nil {
		if errors.Is(err, api.ErrPayload) {
			return ctx.Status(400).SendString("email already exists")
		}
		return ctx.Status(500).SendString("internal server error")
	}
	return ctx.Status(200).JSON(fiber.Map{
		"status": "ok",
	})
}

func Login(ctx *fiber.Ctx) error {
	req := new(User)
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	token, err := Service.Login(*req)

	if err != nil {
		if errors.Is(err, api.ErrPayload) {
			return ctx.Status(400).SendString("incorrect password or email")
		}
		return ctx.Status(500).SendString("internal server error")
	}
	return ctx.Status(200).JSON(fiber.Map{
		"token": token,
	})
}
