package main

import (
	"context"
	"fmt"
	"github.com/edgedb/edgedb-go"
	"github.com/google/uuid"
	"log"
	"strings"
)

type JobTitle struct {
	edgedb.Optional
	Id   uuid.UUID `edgedb:"id"`
	Name string    `edgedb:"name"`
}

type Department struct {
	Id        uuid.UUID  `edgedb:"id"`
	Name      string     `edgedb:"name"`
	Employees []Employee `edgedb:"employees"`
}

type Employee struct {
	Id        uuid.UUID `edgedb:"id"`
	FirstName string    `edgedb:"first_name"`
	LastName  string    `edgedb:"last_name"`
	BirthDate string    `edgedb:"birthday"`

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

	var employees []Employee
	err = client.Query(ctx, "SELECT Employee{first_name, last_name, job_title: {name}, departements:{name}}", &employees)
	if err != nil {
		log.Fatal(err)
	}

	for _, employee := range employees {
		departmentNames := make([]string, len(employee.Departments))
		for i, department := range employee.Departments {
			departmentNames[i] = department.Name
		}
		departmentsConcatenated := strings.Join(departmentNames, ", ")
		log.Printf("Employee: %s %s\n", employee.FirstName, employee.LastName)
		log.Printf("Job Title: %s\n", employee.JobTitle.Name)
		log.Printf("Departments: %s\n", departmentsConcatenated)
		fmt.Println("------")
	}
}
