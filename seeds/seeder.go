package seeds

import (
	"log"
	"reflect"

	"gorm.io/gorm"
)

func Execute(db *gorm.DB, seedMethodNames ...string) {
	s := Seed{db}

	seedType := reflect.TypeOf(s)

	if len(seedMethodNames) == 0 {
		log.Println("Running all seeder...")
		for i := 0; i < seedType.NumMethod(); i++ {
			method := seedType.Method(i)
			seed(s, method.Name)
		}
	}

	for _, item := range seedMethodNames {
		seed(s, item)
	}
}
func seed(s Seed, seedMethodName string) {
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}
	log.Println("Seeding", seedMethodName, "...")
	m.Call(nil)
	log.Println("Seed", seedMethodName, "succedd")
}
