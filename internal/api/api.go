package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kodsurfer/go-happy-birthday/internal/service"
	"strconv"
)

// EmployeeAPI provides HTTP handlers for Employee operations.
type EmployeeAPI struct {
	service *service.EmployeeService
}

// NewEmployeeAPI creates a new EmployeeAPI.
func NewEmployeeAPI(service *service.EmployeeService) *EmployeeAPI {
	return &EmployeeAPI{service: service}
}

// RegisterRoutes registers the routes for EmployeeAPI.
func (api *EmployeeAPI) RegisterRoutes(router *gin.Engine) {
	router.GET("/employees", api.GetAllEmployees)
	router.GET("/employees/:id", api.GetEmployeeByID)
	// Add other routes as needed.
}

// GetAllEmployees returns all employees.
func (api *EmployeeAPI) GetAllEmployees(c *gin.Context) {
	employees, err := api.service.GetAllEmployees()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, employees)
}

// GetEmployeeByID returns an employee by ID.
func (api *EmployeeAPI) GetEmployeeByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	employee, err := api.service.GetEmployeeByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, employee)
}

// Similar methods for SubscriptionAPI can be implemented.
