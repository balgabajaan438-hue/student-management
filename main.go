
package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Group string `json:"group"`
}

var students = []Student{
	{ID: 1, Name: "Ayan", Age: 18, Group: "IT-23"},
	{ID: 2, Name: "Arman", Age: 19, Group: "IT-24"},
}

var nextID = 3

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Student Management System!")
	})

	// Барлық студенттер
	e.GET("/students", func(c echo.Context) error {
		return c.JSON(http.StatusOK, students)
	})

	// Жаңа студент қосу
	e.POST("/students", func(c echo.Context) error {
		var student Student

		if err := c.Bind(&student); err != nil {
			return err
		}

		student.ID = nextID
		nextID++

		students = append(students, student)

		return c.JSON(http.StatusCreated, student)
	})

	// Студентті табу
	e.GET("/students/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		for _, s := range students {
			if s.ID == id {
				return c.JSON(http.StatusOK, s)
			}
		}

		return c.JSON(http.StatusNotFound, "Student not found")
	})

	// Студентті өзгерту
	e.PUT("/students/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		var updated Student
		if err := c.Bind(&updated); err != nil {
			return err
		}

		for i, s := range students {
			if s.ID == id {
				updated.ID = id
				students[i] = updated
				return c.JSON(http.StatusOK, updated)
			}
		}

		return c.JSON(http.StatusNotFound, "Student not found")
	})

	// Студентті өшіру
	e.DELETE("/students/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		for i, s := range students {
			if s.ID == id {
				students = append(students[:i], students[i+1:]...)
				return c.JSON(http.StatusOK, "Student deleted")
			}
		}

		return c.JSON(http.StatusNotFound, "Student not found")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
