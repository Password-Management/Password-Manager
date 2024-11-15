package server

import (
	"fmt"
	"log"
	"password-manager/db"
	"password-manager/handlers"
	"password-manager/queue"
	"password-manager/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func GetDbCheck() {
	done := make(chan struct{})
	go func() {
		fmt.Println("Inside the channel")
		time.Sleep(5 * time.Second)
		db, err := db.NewDbRequest()
		if err != nil {
			log.Fatal("error in creating a DB request")
			return
		}
		_, err = db.InitDB()
		if err != nil {
			log.Println("error in starting the DataBase: ", err)
			return
		}
		close(done)
	}()
}

func Server() {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                              // Specify allowed origin(s)
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH", // Specify allowed methods
	}))
	GetDbCheck()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "200",
			"message": "Server is running smoothly",
		})
	})
	var Log *log.Logger
	go func() {
		time.Sleep(15 * time.Second)
		err := queue.QueueConsumer()
		if err != nil {
			log.Println("error while reading the queue: " + err.Error())
			return
		}
		log.Println("QUEUE IS READY FOR RECEIVING CONNECTION CREATE PRODUCT ENTRY")
	}()
	ser, err := services.NewMasterServiceRequest()
	if err != nil {
		log.Println("master service instance starting failure: " + err.Error())
		return
	}
	userService, err := services.UserServiceRequest()
	if err != nil {
		log.Println("user service instance starting failure: " + err.Error())
		return
	}
	adminService, err := services.NewAdminService()
	if err != nil {
		log.Println("admin service instance starting failure: " + err.Error())
		return
	}
	loginService, err := services.LoginServiceRequest()
	if err != nil {
		log.Println("login service instance starting failure: " + err.Error())
		return
	}
	handlers := handlers.NewHandler(Log).MasterHandler(ser).UserHandler(userService).AdminHandler(adminService).LoginHandler(loginService)
	Routes(app, handlers)
	err = app.Listen(":8000")
	if err != nil {
		log.Print("error in starting the server:", err)
	}
}
