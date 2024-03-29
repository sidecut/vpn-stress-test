package main

import (
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var charset = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

var buffer = make([]byte, 1024/8)

// Send a block of 1024 bits, i.e. 128 bytes
func get1kbBlock(c echo.Context) error {
	for i := 0; i < 1024/8; i++ {
		index := rand.IntN(64)
		buffer[i] = charset[index]
	}
	c.Response().Write(buffer)
	return nil
}

// Send a block of 1024 bytes
func get1kBBlock(c echo.Context) error {
	for i := 0; i < 8; i++ {
		get1kbBlock(c)
	}
	return nil
}

// Send a block of 1048576 bits, i.e. 131072 bytes
func handleGet1MbBlocks(c echo.Context) error {
	blocks, err := strconv.Atoi(c.Param("blocks"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	for i := 0; i < blocks; i++ {
		get1MbBlock(c)
	}
	return nil
}

func get1MbBlock(c echo.Context) error {
	for j := 0; j < 1024/8; j++ {
		get1kBBlock(c)
	}
	return nil
}

// Send a block of 1048576 bytes
func handleGet1MBBlocks(c echo.Context) error {
	blocks, err := strconv.Atoi(c.Param("blocks"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	for i := 0; i < blocks; i++ {
		get1MBBlock(c)
	}
	return nil
}

func get1MBBlock(c echo.Context) error {
	for j := 0; j < 1024; j++ {
		get1kBBlock(c)
	}
	return nil
}

func handleGet1kbBlocks(c echo.Context) error {
	blocks, err := strconv.Atoi(c.Param("blocks"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	for i := 0; i < blocks; i++ {
		get1kbBlock(c)
	}
	return nil
}

func handleGet1kBBlocks(c echo.Context) error {
	blocks, err := strconv.Atoi(c.Param("blocks"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	for i := 0; i < blocks; i++ {
		get1kBBlock(c)
	}
	return nil
}

func handleGetUnitsBlocks(c echo.Context) error {
	blocks, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	units := c.Param("units")
	return getUnitsBlocks(units, blocks, c)

}

func getUnitsBlocks(units string, blocks int, c echo.Context) error {
	var handler echo.HandlerFunc

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
		return echo.ErrBadRequest
	}
	for i := 0; i < blocks; i++ {
		handler(c)
	}
	return nil
}

func handleGetBlocks(c echo.Context) error {
	units := c.Param("units")
	return getUnitsBlocks(units, 1, c)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/api/get1MBBlocks/:blocks", handleGet1MBBlocks)
	e.GET("/api/get1kBBlocks/:blocks", handleGet1kBBlocks)
	e.GET("/api/get1MbBlocks/:blocks", handleGet1MbBlocks)
	e.GET("/api/get1kbBlocks/:blocks", handleGet1kbBlocks)
	e.GET("/api/getBlocks/:units", handleGetBlocks)
	e.GET("/api/getBlocks/:units/:number", handleGetUnitsBlocks)
	e.Logger.Fatal(e.Start(":1323"))
}
