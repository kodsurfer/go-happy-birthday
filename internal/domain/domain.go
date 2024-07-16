package domain

import "time"

// Employee represents an employee in the system.
type Employee struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Birthday time.Time `json:"birthday"`
	Email    string    `json:"email"`
}

// Subscription represents a subscription to a birthday notification.
type Subscription struct {
	ID         int `json:"id"`
	EmployeeID int `json:"employee_id"`
	Subscriber int `json:"subscriber"`
}

// EmployeeRepository defines the interface for employee repository.
type EmployeeRepository interface {
	GetAll() ([]Employee, error)
	GetByID(id int) (*Employee, error)
	Save(employee *Employee) error
	// Add other methods as needed.
}

// SubscriptionRepository defines the interface for subscription repository.
type SubscriptionRepository interface {
	GetByEmployeeID(employeeID int) ([]Subscription, error)
	Subscribe(employeeID, subscriberID int) error
	Unsubscribe(employeeID, subscriberID int) error
	// Add other methods as needed.
}
