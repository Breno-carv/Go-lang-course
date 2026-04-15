package main

import "fmt"

// structs are just key value pairs

type Car struct {
	Make  string
	Model string
	Year  int
	Price float64
}

// You can also Nest types:
type Person struct {
	Name          string
	Age           int
	Car           Car
	FinancialInfo FinancialInfo
}

func newCar(make string, model string, year int, price float64) Car {
	return Car{
		Make:  make,
		Model: model,
		Year:  year,
		Price: price,
	}
}

type FinancialInfo struct {
	Salary        float64
	Debt          float64
	HasActiveDebt bool
}

func (p Person) GetFinancialInfo() FinancialInfo {
	// In a real application, this data would probably come from a database or an API, but for the sake of this example, we will just return some dummy data.
	return FinancialInfo{
		Salary:        p.FinancialInfo.Salary,
		Debt:          p.FinancialInfo.Debt,
		HasActiveDebt: p.FinancialInfo.HasActiveDebt,
	}
}

func newPerson(name string, age int, car Car, financialInfo FinancialInfo) Person {
	return Person{
		Name:          name,
		Age:           age,
		Car:           car,
		FinancialInfo: financialInfo,
	}
}

func main() {

	honda_civic := newCar("Honda", "Civic", 2020, 20000.00)
	john := newPerson("John", 30, honda_civic, FinancialInfo{
		Salary:        50000.00,
		Debt:          10000.00,
		HasActiveDebt: true,
	})

	fmt.Printf("The car is a %d %s %s and it costs $%.2f\n", honda_civic.Year, honda_civic.Make, honda_civic.Model, honda_civic.Price)

	fmt.Printf("%s is %d years old and drives a %d %s %s\n", john.Name, john.Age, john.Car.Year, john.Car.Make, john.Car.Model)

	// we can also initialize a new object/struct with this syntax:
	jane := Person{
		Name: "Jane",
		Age:  28,
		Car: Car{
			Make:  "Toyota",
			Model: "Corolla",
			Year:  2019,
			Price: 18000.00,
		},
	}

	// more familiar, reminds of javascript
	fmt.Printf("%s is %d years old and drives a %d %s %s\n", jane.Name, jane.Age, jane.Car.Year, jane.Car.Make, jane.Car.Model)

	// It is also, obviously, a good practice to verify wether the structs are empty or not, let's say:

	// Let's say my API should return the Person and the FinancialInfo, but some data is still missing, we need to verify
	// that before returning it to the FE, in this example:

	robert := Person{
		Name: "Robert",
		Age:  35,
	}

	if robert.Car.Make == "" || robert.Car.Model == "" || robert.Car.Year == 0 || robert.Car.Price == 0 {
		fmt.Println("The car data is incomplete!")
	}

	if robert.FinancialInfo.Salary == 0 || robert.FinancialInfo.Debt == 0 {
		fmt.Println("The financial info data is incomplete!")
	}

	fmt.Println(robert)

	// There's also anonymous structs:

	anonymousCar := struct {
		Make  string
		Model string
		Year  int
	}{
		Make:  "Ford",
		Model: "Mustang",
		Year:  2021,
	}
	// Honestly, i don't like this kind of syntax at all, maybe in some edge cases it could be useful
	// but in general, i have strong feelings againt it, i mean, just too verbose

	fmt.Printf("The anonymous car is a %d %s %s\n", anonymousCar.Year, anonymousCar.Make, anonymousCar.Model)

	// there's also embedded structs, reminds of inheritance in OOP
	type Employee struct {
		Person   // it receives all fields present in the Person struct, but we can also add new fields to it
		Position string
	}

	// Finally we can create METHODS, they are just functions that have a receiver, in this case, the receiver is the Person struct, but it could be any struct, and we can call it whatever we want, but by convention, we use the first letter of the struct as the receiver name, in this case, we will use "p" for Person.
	james := Employee{
		Person: Person{
			Name: "James",
			Age:  40,
			Car: Car{
				Make:  "BMW",
				Model: "X5",
				Year:  2022,
				Price: 50000.00,
			},
			FinancialInfo: FinancialInfo{
				Salary:        100000.00,
				Debt:          20000.00,
				HasActiveDebt: true,
			},
		},
		Position: "Manager",
	}

	fmt.Println(james.GetFinancialInfo())
}
