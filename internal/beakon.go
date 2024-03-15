package internal

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/gtamasi/beakon/templates"
	"github.com/gtamasi/beakon/templates/components"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func renderView(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func NewBeakon() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Static("/", "assets")

	employees := []components.Employee{
		{
			Id:     0,
			Name:   "John Doe",
			Role:   "CEO",
			Email:  "john.doe@peakon.com",
			Status: "Active",
			Teams:  []string{"Engineering", "Product", "Leadership"},
		},
		{
			Id:     1,
			Name:   "Jane Doe",
			Role:   "CTO",
			Email:  "jane.doe@peakon.com",
			Status: "Active",
			Teams:  []string{"Engineering", "Product", "Leadership"},
		},
		{
			Id:     2,
			Name:   "Harriet Banks",
			Role:   "Manager",
			Email:  "harriet.banks@peakon.com",
			Status: "Active",
			Teams:  []string{"Engineering", "Product"},
		},
		{
			Id:     3,
			Name:   "Gergo Tamasi",
			Role:   "Software Engineer",
			Email:  "gergo.tamasi@peakon.com",
			Status: "Inactive",
			Teams:  []string{"Engineering"},
		},
	}

	e.GET("/login", func(c echo.Context) error {
		return renderView(c, templates.Login("Beakon"))
	})

	e.POST("/login/email", func(c echo.Context) error {
		email := c.FormValue("email")
		// regex for emamil check
		if email == "" {
			return c.NoContent(401)
		}
		// check if the email is valid format with regex
		regex, _ := regexp.Compile(`^[\w-\\.]+@([\w-]+\.)+[\w-]{2,4}`)
		if !regex.MatchString(email) {
			return renderView(c, components.ErrorLoginEmail("Invalid email!"))
		}

		if strings.Compare(email, "test@test.com") == 0 {
			return renderView(c, components.ErrorLoginEmail("Email is already in use!"))
		}

		return c.NoContent(200)
	})

	e.GET("/dashboard", func(c echo.Context) error {
		return renderView(c, templates.Dashboard("Beakon"))
	})

	e.GET("/dashboard/training-card", func(c echo.Context) error {
		return renderView(c, components.TrainingCard())
	})

	e.GET("/explore", func(c echo.Context) error {
		return renderView(c, templates.Explore())
	})

	e.GET("/admin", func(c echo.Context) error {
		return renderView(c, templates.AdminWithLayout("Beakon"))
	})

	e.GET("/admin/employees", func(c echo.Context) error {
		return renderView(c, components.EmployeeTable(employees))
	})

	e.POST("/employees/search", func(c echo.Context) error {
		search := c.FormValue("search")
		var results []components.Employee
		for _, employee := range employees {
			if strings.Contains(strings.ToLower(employee.Name), strings.ToLower(search)) {
				results = append(results, employee)
			}
		}
		return renderView(c, components.EmployeeTable(results))
	})

	e.DELETE("/employees/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		for i, employee := range employees {
			if employee.Id == id {
				employees = append(employees[:i], employees[i+1:]...)
				break
			}
		}

		return renderView(c, components.EmployeeTable(employees))
	})

	e.Logger.Fatal(e.Start(":3003"))
}
