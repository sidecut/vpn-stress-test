package main

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var charset = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

var buffer = make([]byte, 1024/8)

// Send a block of 1024 bits, i.e. 128 bytes
func get1kbBlock(c *fiber.Ctx) error {
	for i := 0; i < 1024/8; i++ {
		index := rand.Intn(64)
		buffer[i] = charset[index]
	}
	return c.Send(buffer)
}

// Send a block of 1024 bytes
func get1kBBlock(c *fiber.Ctx) error {
	for i := 0; i < 8; i++ {
		err := get1kbBlock(c)
		if err != nil {
			return err
		}
	}
	return nil
}

// Send a block of 1048576 bits, i.e. 131072 bytes
func handleGet1MbBlocks(c *fiber.Ctx) error {
	blocks, err := strconv.Atoi(c.Params("blocks"))
	if err != nil {
		return err
	}
	for i := 0; i < blocks; i++ {
		err := get1MbBlock(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func get1MbBlock(c *fiber.Ctx) error {
	for j := 0; j < 1024/8; j++ {
		err := get1kBBlock(c)
		if err != nil {
			return err
		}
	}
	return nil
}

// Send a block of 1048576 bytes
func handleGet1MBBlocks(c *fiber.Ctx) error {
	blocks, err := strconv.Atoi(c.Params("blocks"))
	if err != nil {
		return err
	}
	for i := 0; i < blocks; i++ {
		err := get1MBBlock(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func get1MBBlock(c *fiber.Ctx) error {
	for j := 0; j < 1024; j++ {
		err := get1kBBlock(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func handleGet1kbBlocks(c *fiber.Ctx) error {
	blocks, err := strconv.Atoi(c.Params("blocks"))
	if err != nil {
		return err
	}
	for i := 0; i < blocks; i++ {
		err := get1kbBlock(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func handleGet1kBBlocks(c *fiber.Ctx) error {
	blocks, err := strconv.Atoi(c.Params("blocks"))
	if err != nil {
		return err
	}
	for i := 0; i < blocks; i++ {
		err := get1kBBlock(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func handleGetUnitsBlocks(c *fiber.Ctx) error {
	blocks, err := strconv.Atoi(c.Params("number"))
	if err != nil {
		return err
	}
	units := c.Params("units")
	return getUnitsBlocks(units, blocks, c)
}

func getUnitsBlocks(units string, blocks int, c *fiber.Ctx) error {
	var handler fiber.Handler

	switch units {
	case "kb":
		handler = get1kbBlock
	case "kB":
		handler = get1kBBlock
	case "Mb":
		handler = get1MbBlock
	case "MB":
		handler = get1MBBlock
	default:
		return fiber.ErrBadRequest
	}
	for i := 0; i < blocks; i++ {
		handler(c)
	}
	return nil
}

func handleGetBlocks(c *fiber.Ctx) error {
	units := c.Params("units")
	return getUnitsBlocks(units, 1, c)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Hello, World!")
	})
	app.Get("/api/get1MBBlocks/:blocks", handleGet1MBBlocks)
	app.Get("/api/get1kBBlocks/:blocks", handleGet1kBBlocks)
	app.Get("/api/get1MbBlocks/:blocks", handleGet1MbBlocks)
	app.Get("/api/get1kbBlocks/:blocks", handleGet1kbBlocks)
	app.Get("/api/getBlocks/:units", handleGetBlocks)
	app.Get("/api/getBlocks/:units/:number", handleGetUnitsBlocks)
	log.Fatal(app.Listen(":1323"))
}
