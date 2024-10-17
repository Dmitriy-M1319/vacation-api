package models

type Employee struct {
	ID         uint64 `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Patronymic string `json:"patronymic"`
}

type Vacation struct {
	ID         uint64 `json:"id"`
	EmployeeID uint64 `json:"empId"`
	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
	DaysCount  uint64 `json:"daysCount"`
}

type VacationNorm struct {
	ID             uint64 `json:"id"`
	Month          string `json:"month"`
	VacationsCount uint64 `json:"vacationsCount"`
}
