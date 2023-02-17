package app

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"os"
	"simpleApi/config"
	"simpleApi/pkg/govalidator"
)

func Run() {
	validator := govalidator.New()

	log := zerolog.New(os.Stdout)

	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	err = validator.Validate(context.Background(), cfg)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	log = log.Level(zerolog.Level(*cfg.Logger.Level)).With().Timestamp().Logger()

	defer log.Info().Msg("Application has been shut down")

	log.Debug().Msg("Loaded configuration")

	app := fiber.New()

	app.Get("/api/create_user", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Ты создал юзера Congratulations! %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	})

	err = app.Listen(":3000")
	if err != nil {
		log.Fatal().Err(err)
	}
}
