package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

var (
	Version    string
	GitCommit  string
	port       string
	configFile string
	ConfigData Config
)

func init() {
	log.Println("Init")

	port = GetEnv("reggie_port", "8080")
	configFile := GetEnv("reggie_config", "config.json")

	jsonFile, err := os.Open(configFile)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("Successfully Loaded config.json")

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &ConfigData)
}

func main() {

	log.Println("Starting")
	log.Println("Version:", Version)
	log.Println("GitCommit:", GitCommit)

	app := fiber.New()

	app.Get("/.well-known/terraform.json", func(c *fiber.Ctx) error {
		log.Println("GET /.well-known/terraform.json")
		c.Accepts("application/json")
		c.Set("content-type", "application/json; charset=utf-8")
		return c.JSON(fiber.Map{"providers.v1": "/terraform/providers/v1/"})

	})

	app.Get("/terraform/providers/v1/:namespace/:name/versions", func(c *fiber.Ctx) error {
		req := fmt.Sprintf("GET /terraform/providers/v1/%s/%s/versions", c.Params("namespace"), c.Params("name"))
		log.Println(req)

		c.Accepts("application/json")
		c.Set("content-type", "application/json; charset=utf-8")

		for _, v := range ConfigData {
			if v.Name == string(c.Params("namespace")+`/`+c.Params("name")) {
				log.Println(v.Name)

				var versions []VersionResp

				for _, v2 := range v.Versions {
					ver := VersionResp{
						Version:   v2,
						Protocols: v.Protocols,
						Platforms: v.Platforms,
					}
					versions = append(versions, ver)
				}

				return c.JSON(Versions{Versions: versions})
			}
		}

		return c.SendString(`getting versions`)
	})

	app.Get("/terraform/providers/v1/:namespace/:name/:version/download/:os/:arch", func(c *fiber.Ctx) error {
		req := fmt.Sprintf("GET /terraform/providers/v1/%s/%s/%s/download/%s/%s", c.Params("namespace"), c.Params("name"), c.Params("version"), c.Params("os"), c.Params("arch"))
		log.Println(req)

		c.Accepts("application/json")
		c.Set("content-type", "application/json; charset=utf-8")

		for _, v := range ConfigData {
			if v.Name == string(c.Params("namespace")+`/`+c.Params("name")) {

				shasum := GetShasum(replaceVars(v.ShasumsURL, c.Params("namespace"), c.Params("name"), c.Params("version"), c.Params("os"), c.Params("arch")), c.Params("os"), c.Params("arch"))

				resp := DownloadResp{
					Protocols:           v.Protocols,
					Os:                  c.Params("os"),
					Arch:                c.Params("arch"),
					Filename:            replaceVars(v.Filename, c.Params("namespace"), c.Params("name"), c.Params("version"), c.Params("os"), c.Params("arch")),
					DownloadURL:         replaceVars(v.DownloadURL, c.Params("namespace"), c.Params("name"), c.Params("version"), c.Params("os"), c.Params("arch")),
					ShasumsURL:          replaceVars(v.ShasumsURL, c.Params("namespace"), c.Params("name"), c.Params("version"), c.Params("os"), c.Params("arch")),
					ShasumsSignatureURL: replaceVars(v.ShasumsSignatureURL, c.Params("namespace"), c.Params("name"), c.Params("version"), c.Params("os"), c.Params("arch")),
					Shasum:              shasum,
					SigningKeys:         v.SigningKeys,
				}

				return c.JSON(resp)
			}
		}

		return c.JSON(&fiber.Map{
			"rrror":   "Not Found",
			"request": req,
		})
	})

	app.Get("/", func(c *fiber.Ctx) error {
		log.Println("GET /")
		c.Accepts("application/json")
		c.Set("content-type", "application/json; charset=utf-8")
		return c.JSON(fiber.Map{"providers.v1": "/terraform/providers/v1/"})
	})

	app.Listen(string(`:` + port))
}


