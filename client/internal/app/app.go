package app

import (
	"clientTagesImages/config"
	"clientTagesImages/internal/domain/dto"
	"clientTagesImages/internal/domain/repository"
	"clientTagesImages/internal/domain/service"
	"clientTagesImages/pkg/govalidator"
	"clientTagesImages/pkg/reqvalidator"
	"clientTagesImages/pkg/storage/postgres"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const maxUploadSize = 2 * 1024 * 1024 // 2 MB
const uploadPath = "./tmp"

func Run() {
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)

	validator := govalidator.New()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = validator.Validate(context.Background(), cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}

	app := fiber.New()

	psqlDB, err := postgres.InitPsqlDB(context.Background(), cfg)
	if err != nil {
		log.Fatal("PostgreSQL init error: %s", err)
	} else {
		log.Infof("PostgreSQL connected, status: %#v", psqlDB.Stats())
	}
	defer func(psqlDB *sqlx.DB) {
		err = psqlDB.Close()
		if err != nil {
			log.Infof(err.Error())
		} else {
			log.Info("PostgreSQL closed properly")
		}
	}(psqlDB)
	UserRepo := repository.NewUserRepo(psqlDB)
	userService := service.NewUserService(UserRepo, *validator)

	app.Post("/api/register", func(c *fiber.Ctx) error {
		var params dto.RegisterUser
		if err := reqvalidator.ReadRequest(c, &params); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		err = userService.Create(c.Context(), params)
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
		}
		return c.SendStatus(fiber.StatusOK)
	})

	app.Post("/api/login", func(c *fiber.Ctx) error {
		var params dto.RegisterUser
		if err := reqvalidator.ReadRequest(c, &params); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		err = userService.Create(c.Context(), params)
		if err != nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusOK).JSON("fuck-you")
	})
	app.Post("/api/upload_image", func(c *fiber.Ctx) error {
		var params dto.UploadImage
		var file *multipart.FileHeader
		multipart, err := c.MultipartForm()
		if err != nil {
			return err
		}
		openedFile, err := file.Open()
		bytes := make([]byte, 10000000)
		openedFile.ReadAt(bytes, 0)

		detectedFileType := http.DetectContentType(bytes)
		switch detectedFileType {
		case "image/jpeg", "image/jpg":
		case "image/gif", "image/png":
		case "application/pdf":
			break
		default:
			return err
		}
		fileEndings, err := mime.ExtensionsByType(detectedFileType)
		if err != nil {
			return err
		}
		fileheader := multipart.File["file"]
		if fileheader != nil {
			file = fileheader[0]
		}
		params.ImageName = c.FormValue("imageName")

		newPath := filepath.Join(uploadPath, file.Filename+fileEndings[0])

		newFile, err := os.Create(newPath)
		if err != nil {
			return err
		}
		defer newFile.Close()
		if _, err := newFile.Write(bytes); err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON("файл успешно загружен")
	})

	app.Post("/api/download_image", func(c *fiber.Ctx) error {
		return err
	})

	app.Post("/api/logout", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Ты создал юзера Congratulations! %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	})

	err = app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
