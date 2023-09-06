package main

import (
	"context"
	"log"
	"time"

	"github.com/edgedb/edgedb-go"
	"github.com/gin-gonic/gin"
)

type JobTitle struct {
	edgedb.Optional
	Id   edgedb.OptionalUUID `json:"id,omitempty" edgedb:"id"`
	Name string              `edgedb:"name"`
}

type Department struct {
	Id        edgedb.OptionalUUID `edgedb:"id"`
	Name      string              `edgedb:"name"`
	Employees []Employee          `edgedb:"employees"`
}

type Employee struct {
	Id        edgedb.OptionalUUID      `edgedb:"id"`
	FirstName string                   `edgedb:"first_name"`
	LastName  string                   `edgedb:"last_name"`
	BirthDate edgedb.OptionalLocalDate `edgedb:"birthday"`

	JobTitle    JobTitle     `edgedb:"job_title"`
	Departments []Department `edgedb:"departements"`
}

func main() {
	ctx := context.Background()

	client, err := edgedb.CreateClientDSN(ctx, "edgedb://edgedb:password@127.0.0.1:5656/super-carnival", edgedb.Options{
		TLSOptions: edgedb.TLSOptions{
			SecurityMode: edgedb.TLSModeInsecure,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		currentTime := time.Now()
		if c.GetHeader("HX-Request") == "true" {
			c.HTML(200, "IndexContent", gin.H{
				"title": "Super Carnival",
				"time":  currentTime.Format("2006.01.02 15:04:05"),
			})
		} else {
			c.HTML(200, "index.html", gin.H{
				"title":       "Super Carnival",
				"time":        currentTime.Format("2006.01.02 15:04:05"),
				"currentPage": "employees",
			})
		}
	})

	r.GET("/job-titles", func(c *gin.Context) {
		currentTime := time.Now()
		if c.GetHeader("HX-Request") == "true" {
			c.HTML(200, "IndexContent", gin.H{
				"title": "Super Carnival",
				"time":  currentTime.Format("2006.01.02 15:04:05"),
			})
		} else {
			c.HTML(200, "index.html", gin.H{
				"title":       "Super Carnival",
				"time":        currentTime.Format("2006.01.02 15:04:05"),
				"currentPage": "job-title",
			})
		}
	})

	r.GET("/departments", func(c *gin.Context) {
		currentTime := time.Now()
		if c.GetHeader("HX-Request") == "true" {
			c.HTML(200, "IndexContent", gin.H{
				"title": "Super Carnival",
				"time":  currentTime.Format("2006.01.02 15:04:05"),
			})
		} else {
			c.HTML(200, "index.html", gin.H{
				"title":       "Super Carnival",
				"time":        currentTime.Format("2006.01.02 15:04:05"),
				"currentPage": "departments",
			})
		}
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/employees", func(c *gin.Context) {
		var employees []Employee
		err = client.Query(ctx, `
			SELECT Employee { 
				**,
				departements: {
					*
				}
			}`, &employees)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(200, employees)
	})

	r.Run()
}
