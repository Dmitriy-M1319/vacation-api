package models

type Employee struct {
	ID         uint64 `json:"id" db:"id"`
	FirstName  string `json:"firstName" db:"first_name"`
	LastName   string `json:"lastName" db:"last_name"`
	Patronymic string `json:"patronymic" db:"patronymic"`
}

type Vacation struct {
	ID         uint64 `json:"id" db:"id"`
	EmployeeID uint64 `json:"empId" db:"emp_id"`
	StartDate  string `json:"startDate" db:"start_date"`
	EndDate    string `json:"endDate" db:"end_date"`
	DaysCount  uint64 `json:"daysCount" db:"days_count"`
}

type VacationNorm struct {
	ID             uint64 `json:"id" db:"id"`
	Month          string `json:"month" db:"month"`
	VacationsCount uint64 `json:"vacationsCount" db:"vacations_count"`
}
