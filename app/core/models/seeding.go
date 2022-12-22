package models

import (
	"encoding/csv"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
	"jatis-test/internal/config"
	"os"
	"strconv"
	"time"
)

func Seeding(data string) {
	if data == "customers" {
		seedCustomers()
	} else if data == "employees" {
		seedEmployees()
	} else if data == "products" {
		seedProducts()
	} else if data == "shipping" {
		seedShipping()
	} else if data == "orders" {
		seedOrders()
	} else if data == "order_detail" {
		seedOrderDetails()
	} else if data == "all" {
		fmt.Println("seed all data")
		seedCustomers()
		seedEmployees()
		seedProducts()
		seedShipping()
		seedOrders()
		seedOrderDetails()
	}
}

func seedCustomers() {
	// open file
	f, err := os.Open("assets/customers.csv")
	if err != nil {
		logrus.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		logrus.Fatal(err)
	}

	// convert records to array of structs
	customers := make([]Customers, 0)
	for i, line := range data {
		if i > 0 {
			customerID, _ := strconv.Atoi(line[0])
			dataCustomers := Customers{
				CustomerID:          customerID,
				CompanyName:         line[1],
				FirstName:           line[2],
				LastName:            line[3],
				BillingAddress:      line[4],
				City:                line[5],
				StateOrProvince:     line[6],
				ZipCode:             line[7],
				Email:               line[8],
				CompanyWebsite:      line[9],
				PhoneNumber:         line[10],
				FaxNumber:           line[11],
				ShipAddress:         line[12],
				ShipCity:            line[13],
				ShipStateOrProvince: line[14],
				ShipZipCode:         line[15],
				ShipPhoneNumber:     line[16],
			}
			customers = append(customers, dataCustomers)
		}
	}
	config.GetInstancePostgresDb().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(customers, 100)

	// print the array
	fmt.Println("Create Customers success")

}

func seedEmployees() {
	// open file
	f, err := os.Open("assets/employees.csv")
	if err != nil {
		logrus.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		logrus.Fatal(err)
	}

	// convert records to array of structs
	employees := make([]Employees, 0)
	for i, line := range data {
		if i > 0 {
			employeeID, _ := strconv.Atoi(line[0])
			dataEmployee := Employees{
				EmployeeID: employeeID,
				FirstName:  line[1],
				LastName:   line[2],
				Title:      line[3],
				WorkPhone:  line[4],
			}
			employees = append(employees, dataEmployee)
		}
	}
	config.GetInstancePostgresDb().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(employees, 100)

	// print the array
	fmt.Println("Create Employees success")
}

func seedProducts() {
	// open file
	f, err := os.Open("assets/products.csv")
	if err != nil {
		logrus.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		logrus.Fatal(err)
	}

	// convert records to array of structs
	products := make([]Products, 0)
	for i, line := range data {
		if i > 0 {
			productID, _ := strconv.Atoi(line[0])
			unitPrice, _ := strconv.Atoi(line[2])
			inStock, _ := strconv.Atoi(line[3])
			dataProduct := Products{
				ProductID:   productID,
				ProductName: line[1],
				UnitPrice:   float64(unitPrice),
				InStock:     inStock,
			}
			products = append(products, dataProduct)
		}
	}
	config.GetInstancePostgresDb().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(products, 100)

	// print the array
	fmt.Println("Create Products success")
}

func seedShipping() {
	// open file
	f, err := os.Open("assets/shipping.csv")
	if err != nil {
		logrus.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		logrus.Fatal(err)
	}

	// convert records to array of structs
	shippings := make([]ShippingMethods, 0)
	for i, line := range data {
		if i > 0 {
			shippingID, _ := strconv.Atoi(line[0])
			dataShipping := ShippingMethods{
				ShippingMethodID: shippingID,
				ShippingMethod:   line[1],
			}
			shippings = append(shippings, dataShipping)
		}
	}
	config.GetInstancePostgresDb().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(shippings, 100)

	// print the array
	fmt.Println("Create Shipping Method success")
}

func seedOrders() {
	// open file
	f, err := os.Open("assets/orders.csv")
	if err != nil {
		logrus.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		logrus.Fatal(err)
	}

	// convert records to array of structs
	orders := make([]Orders, 0)
	for i, line := range data {
		if i > 0 {
			orderID, _ := strconv.Atoi(line[0])
			customerID, _ := strconv.Atoi(line[1])
			employeeID, _ := strconv.Atoi(line[2])
			orderDate, _ := time.Parse("2006-01-02", line[3])
			shipDate, _ := time.Parse("2006-01-02", line[5])
			shippingID, _ := strconv.Atoi(line[6])
			charge, _ := strconv.Atoi(line[7])
			taxes, _ := strconv.Atoi(line[8])
			payment, _ := strconv.Atoi(line[9])
			dataOrder := Orders{
				OrderID:             orderID,
				CustomerID:          customerID,
				EmployeeID:          employeeID,
				OrderDate:           orderDate,
				PurchaseOrderNumber: line[4],
				ShipDate:            shipDate,
				ShippingMethodID:    shippingID,
				FreightCharge:       float64(charge),
				Taxes:               float64(taxes),
				PaymentReceived:     payment,
				Comment:             line[10],
			}
			orders = append(orders, dataOrder)
		}
	}
	config.GetInstancePostgresDb().Omit(clause.Associations).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(orders, 100)

	// print the array
	fmt.Println("Create Orders success")
}

func seedOrderDetails() {
	// open file
	f, err := os.Open("assets/order_details.csv")
	if err != nil {
		logrus.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		logrus.Fatal(err)
	}

	// convert records to array of structs
	orderDetails := make([]OrderDetails, 0)
	for i, line := range data {
		if i > 0 {
			orderDetailID, _ := strconv.Atoi(line[0])
			orderID, _ := strconv.Atoi(line[1])
			productID, _ := strconv.Atoi(line[2])
			quantity, _ := strconv.Atoi(line[3])
			unitPrice, _ := strconv.Atoi(line[4])
			discount, _ := strconv.Atoi(line[5])
			dataOrderDetail := OrderDetails{
				OrderDetailID: orderDetailID,
				OrderID:       orderID,
				ProductID:     productID,
				Quantity:      quantity,
				UnitPrice:     float64(unitPrice),
				Discount:      float64(discount),
			}
			orderDetails = append(orderDetails, dataOrderDetail)
		}
	}
	config.GetInstancePostgresDb().Omit(clause.Associations).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(orderDetails, 100)

	// print the array
	fmt.Println("Create Order Details success")
}
