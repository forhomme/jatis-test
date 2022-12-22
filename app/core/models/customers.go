package models

type Customers struct {
	CustomerID          int    `json:"customer_id" gorm:"primaryKey"`
	CompanyName         string `json:"company_name"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	BillingAddress      string `json:"billing_address"`
	City                string `json:"city"`
	StateOrProvince     string `json:"state_or_province"`
	ZipCode             string `json:"zip_code"`
	Email               string `json:"email"`
	CompanyWebsite      string `json:"company_website"`
	PhoneNumber         string `json:"phone_number"`
	FaxNumber           string `json:"fax_number"`
	ShipAddress         string `json:"ship_address"`
	ShipCity            string `json:"ship_city"`
	ShipStateOrProvince string `json:"ship_state_or_province"`
	ShipZipCode         string `json:"ship_zip_code"`
	ShipPhoneNumber     string `json:"ship_phone_number"`
}

func (Customers) TableName() string {
	return "customers"
}

func (Customers) ModuleName() string {
	return "customers"
}
