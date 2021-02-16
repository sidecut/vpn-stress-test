package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var charset = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

var buffer = make([]byte, 1024/8)

// Send a block of 1024 bits, i.e. 128 bytes
func get1kbBlock(c echo.Context) error {
	for i := 0; i < 1024/8; i++ {
		index := rand.Intn(64)
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
		for j := 0; j < 1024/8; j++ {
			get1kBBlock(c)
		}
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
		for j := 0; j < 1024; j++ {
			get1kBBlock(c)
		}
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

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/api/get1MBBlocks/:blocks", handleGet1MBBlocks)
	e.GET("/api/get1kBBlocks/:blocks", handleGet1kBBlocks)
	e.GET("/api/get1MbBlocks/:blocks", handleGet1MbBlocks)
	e.GET("/api/get1kbBlocks/:blocks", handleGet1kbBlocks)
	e.Logger.Fatal(e.Start(":1323"))
}
