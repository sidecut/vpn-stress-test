package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var charset = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

// Send a block of 1024 bits
func get1kbBlock(c echo.Context) error {
	for i := 0; i < 1024/8; i++ {
		index := rand.Intn(64)
		c.Response().Write(charset[index : index+1])
	}
	return nil
}

// Send a block of 1024 characters
func get1kBBlock(c echo.Context) error {
	for i := 0; i < 1024; i++ {
		index := rand.Intn(64)
		c.Response().Write(charset[index : index+1])
	}
	return nil
}

func get1MbBlocks(c echo.Context) error {
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

func get1MBBlocks(c echo.Context) error {
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

func get1kbBlocks(c echo.Context) error {
	blocks, err := strconv.Atoi(c.Param("blocks"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	for i := 0; i < blocks; i++ {
		get1kbBlock(c)
	}
	return nil
}

func get1kBBlocks(c echo.Context) error {
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
	e.GET("/api/get1MBBlocks/:blocks", get1MBBlocks)
	e.GET("/api/get1kBBlocks/:blocks", get1kBBlocks)
	e.GET("/api/get1MbBlocks/:blocks", get1MbBlocks)
	e.GET("/api/get1kbBlocks/:blocks", get1kbBlocks)
	e.Logger.Fatal(e.Start(":1323"))
}
