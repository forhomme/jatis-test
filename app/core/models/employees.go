package models

type Employees struct {
	EmployeeID int    `json:"employee_id" gorm:"primaryKey"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Title      string `json:"title"`
	WorkPhone  string `json:"work_phone"`
}

func (Employees) TableName() string {
	return "employees"
}

func (Employees) ModuleName() string {
	return "employees"
}
