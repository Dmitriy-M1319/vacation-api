package models

type EmployeeDto struct {
	Employee
}

type EmployeeBodyDto struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Patronymic string `json:"patronymic"`
}

type VacationDto struct {
	Vacation
}

type VacationBodyDto struct {
	EmployeeID uint64 `json:"empId"`
	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
	DaysCount  uint64 `json:"daysCount"`
}

type VacationNormDto struct {
	VacationNorm
}

type VacationNormBodyDto struct {
	Month          string `json:"month"`
	VacationsCount uint64 `json:"vacationsCount"`
}
