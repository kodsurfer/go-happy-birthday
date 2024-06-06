package service

import "github.com/yourusername/birthday-service/domain"

// EmployeeService provides operations on Employee.
type EmployeeService struct {
	repo domain.EmployeeRepository
}

// NewEmployeeService creates a new EmployeeService.
func NewEmployeeService(repo domain.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

// GetAllEmployees returns all employees.
func (s *EmployeeService) GetAllEmployees() ([]domain.Employee, error) {
	return s.repo.GetAll()
}

// GetEmployeeByID returns an employee by ID.
func (s *EmployeeService) GetEmployeeByID(id int) (*domain.Employee, error) {
	return s.repo.GetByID(id)
}

// SaveEmployee saves an employee.
func (s *EmployeeService) SaveEmployee(employee *domain.Employee) error {
	return s.repo.Save(employee)
}

// SubscriptionService provides operations on Subscription.
type SubscriptionService struct {
	repo domain.SubscriptionRepository
}

// NewSubscriptionService creates a new SubscriptionService.
func NewSubscriptionService(repo domain.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

// GetSubscriptionsByEmployeeID returns subscriptions for an employee.
func (s *SubscriptionService) GetSubscriptionsByEmployeeID(employeeID int) ([]domain.Subscription, error) {
	return s.repo.GetByEmployeeID(employeeID)
}

// Subscribe subscribes a user to an employee's birthday notifications.
func (s *SubscriptionService) Subscribe(employeeID, subscriberID int) error {
	return s.repo.Subscribe(employeeID, subscriberID)
}

// Unsubscribe unsubscribes a user from an employee's birthday notifications.
func (s *SubscriptionService) Unsubscribe(employeeID, subscriberID int) error {
	return s.repo.Unsubscribe(employeeID, subscriberID)
}
