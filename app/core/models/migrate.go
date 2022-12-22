package models

import "jatis-test/internal/config"

func Migrate() {
	config.GetInstancePostgresDb().AutoMigrate(
		&Customers{},       // migrate customers database
		&Employees{},       // migrate employees database
		&Products{},        // migrate products database
		&ShippingMethods{}, // migrate shipping methods database
		&Orders{},          // migrate orders database
		&OrderDetails{},    // migrate order details database
	)
}
