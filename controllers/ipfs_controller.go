package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	shell "github.com/ipfs/go-ipfs-api"
	"log"
)

func PostUploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to retrieve file from request",
		})
	}

	sh := shell.NewShell("localhost:5001")

	f, err := file.Open()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to retrieve file from request",
		})
	}
	defer f.Close()

	cid, err := sh.Add(f)
	fmt.Println(cid)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to upload file to IPFS",
		})
	}

	// example https://ipfs.io/ipfs/QmPoN2mvHWXj5TZzJvuDafUKRFBwarp9wtmx5QZ5LW8TCT
	return c.JSON(fiber.Map{
		"Status":  fiber.StatusOK,
		"Message": "File uploaded to IPFS",
		"Data": fiber.Map{
			"url": fmt.Sprintf("https://ipfs.io/ipfs/%s", cid),
			"cid": cid, // secure
		},
	})
}
