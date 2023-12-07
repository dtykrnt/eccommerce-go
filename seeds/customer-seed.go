package seeds

import (
	"golang-basic/modules/customers"

	"github.com/go-faker/faker/v4"
	"gorm.io/gorm"
)

// Seed type
type Seed struct {
	db *gorm.DB
}

func (s Seed) CustomerSeed() {
	for i := 0; i <= 10; i++ {
		cust := customers.Customer{Name: faker.Name(), Email: faker.Email()}
		s.db.Create(&cust)
	}
}
