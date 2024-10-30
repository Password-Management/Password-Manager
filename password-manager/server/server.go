package server

import (
	"fmt"
	"log"
	dallayer "password-manager/dalLayer"
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
		resp, err := db.InitDB()
		if err != nil {
			log.Println("error in starting the DataBase: ", err)
		}
		fmt.Println("before If condition")
		if resp != nil {
			fmt.Println("Inside the RESP BLOCK !!!!!")
			log.Println("THE DATABASE IS RUNNING")
			dal, err := dallayer.NewMasterDalRequest()
			if err != nil {
				log.Println("error in checking the BlockChain (setting instance):", err)
			}
			resp, err := dal.FindAll()
			if err != nil {
				log.Println("error in checking the BlockChain (findALL):", err)
			}
			if len(resp) == 0 {
				log.Println("Has no entry")
				master, _ := services.NewMasterServiceRequest()
				err := master.Create()
				if err != nil {
					log.Println(err, "Exiting the code")
					return
				}

			} else {
				log.Println("DataBase already has the master entry")
			}
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
	go func() {
		time.Sleep(10*time.Second)
		queue.QueueConsumer()
	}()
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "200",
			"message": "Server is running smoothly",
		})
	})
	var Log *log.Logger
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
	handlers := handlers.NewHandler(Log).MasterHandler(ser).UserHandler(userService)
	Routes(app, handlers)
	err = app.Listen(":8000")
	if err != nil {
		log.Print("error in starting the server:", err)
	}
}
