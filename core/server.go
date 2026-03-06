package core

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml: "server"`

	AllowedExtensions []string `yaml:"allowed_ex"`
}

var cfg *Config

func loadcfg() error {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return err
	}
	cfg = &Config{}

	return yaml.Unmarshal(data, cfg)
}

func listBooks(extensions []string) (error, int, []string) {
	validExts := make(map[string]bool)
	for _, ext := range extensions {
		validExts[ext] = true
	}

	entries, err := os.ReadDir("books")
	if err != nil {
		return err, 0, nil
	}

	var found []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		fileExt := filepath.Ext(entry.Name())
		if validExts[fileExt] {
			found = append(found, entry.Name())
		}
	}

	return nil, len(found), found
}

func Server() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	if err := loadcfg(); err != nil {
		fmt.Println("[SERVER]: Error, can`t read config")
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "pong",
		})
	})

	r.GET("/books", func(c *gin.Context) {
		allowed_ex := cfg.AllowedExtensions
		err, _, files := listBooks(allowed_ex)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		type BookItem struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}

		var response []BookItem
		for _, f := range files {
			response = append(response, BookItem{
				Name: f,
				Url:  "/books/" + f,
			})
		}

		c.JSON(http.StatusOK, response)
	})

	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	r.GET("/config.yaml", func(c *gin.Context) {
		c.File("config.yaml")
	})

	r.GET("/books/:filename", func(c *gin.Context) {
		nm := c.Param("filename")
		c.FileAttachment("./books/"+nm, nm)
	})

	r.POST("/books/upload", func(c *gin.Context) {
		if err := os.MkdirAll("books", os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directory"})
			return
		}
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(404, gin.H{
				"Error": "File not found",
			})
			return
		}
		filename := filepath.Base(file.Filename)
		dst := filepath.Join("books", filename)
		ext := filepath.Ext(file.Filename)
		allow := false
		for _, allowed := range cfg.AllowedExtensions {
			if ext == allowed {
				allow = true
				break
			}
		}
		if !allow {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Unknown extenstion",
			})
			return
		}

		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(500, gin.H{
				"error": "Couldn`t save the file",
			})
		}
		c.JSON(200, gin.H{
			"status": "True",
		})
	})

	r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}
