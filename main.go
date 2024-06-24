package main

import (
	"GoGo_Database/database"
	"fmt"

	//  "io/ioutil"
	// "encoding/json"

	"go.mongodb.org/mongo-driver/bson"
)

type Address struct {
	City    string `bson:"city"`
	State   string `bson:"state"`
	Country string `bson:"country"`
	Pincode string `bson:"pincode"`
}

type User struct {
	Name    string  `bson:"name"`
	Age     uint32  `bson:"age"`
	Company string  `bson:"company"`
	Contact string  `bson:"contact"`
	Address Address `bson:"inline"`
}

func main() {
	dir := "./"
	db, err := database.New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []User{
		{"John", 23, "23344333", "Myrl Tech", Address{"Bangalore", "Karnataka", "India", "410013"}},
		{"Emily", 30, "111223344", "Tech Solutions Ltd.", Address{"New York", "NY", "USA", "10001"}},
		{"Michael", 28, "987654321", "Data Systems Inc.", Address{"San Francisco", "CA", "USA", "94105"}},
		{"Sophia", 25, "555666777", "Global Solutions", Address{"London", "", "UK", "11111"}},
		{"David", 32, "444555666", "Tech Innovations", Address{"Berlin", "", "Germany", "10115"}},
		{"Linda", 27, "777888999", "Digital Creations", Address{"Sydney", "NSW", "Australia", "2000"}},
		{"Daniel", 29, "123456789", "Innovate IT", Address{"Toronto", "ON", "Canada", "99999"}},
		{"Sarah", 26, "999000111", "Future Tech Co.", Address{"Tokyo", "", "Japan", "10001"}},
		{"Mark", 31, "222333444", "Global Ventures", Address{"Dubai", "", "UAE", "12345"}},
		{"Emma", 24, "666777888", "Smart Solutions", Address{"Paris", "", "France", "75001"}},
		{"James", 33, "888999000", "Innovative Technologies", Address{"Moscow", "", "Russia", "10100"}},
		{"Olivia", 29, "1122334455", "Tech Experts", Address{"Singapore", "", "Singapore", "49178"}},
		{"Alexander", 28, "3344556677", "Digital Solutions Inc.", Address{"Seoul", "", "South Korea", "100011"}},
		{"Grace", 34, "777777777", "Creative Minds", Address{"Los Angeles", "CA", "USA", "90001"}},
		{"William", 31, "555555555", "Tech Innovate", Address{"London", "", "UK", "999999"}},
		{"Sophie", 27, "333333333", "Data Systems", Address{"Berlin", "", "Germany", "10117"}},
		{"Benjamin", 29, "666666666", "Digital Design", Address{"Sydney", "NSW", "Australia", "2000"}},
		{"Isabella", 32, "999999999", "Global Tech", Address{"Toronto", "ON", "Canada", "999999"}},
		{"Jacob", 25, "111111111", "Tech Solutions", Address{"Tokyo", "", "Japan", "100001"}},
		{"Ava", 26, "444444444", "Future Innovations", Address{"Dubai", "", "UAE", "12345"}},
		{"Matthew", 32, "888888888", "Smart Systems", Address{"Paris", "", "France", "75001"}},
		{"Amelia", 23, "222222222", "Innovative Tech", Address{"Moscow", "", "Russia", "101000"}},
		{"Lucas", 28, "777777777", "Tech Ventures", Address{"Singapore", "", "Singapore", "49178"}},
		{"Mia", 29, "555555555", "Digital Solutions", Address{"Seoul", "", "South Korea", "100011"}},
		{"Daniel", 27, "333333333", "Creative Solutions", Address{"Los Angeles", "CA", "USA", "90001"}},
		{"Charlotte", 31, "666666666", "Tech Creations", Address{"London", "", "UK", "999999"}},
		{"Ethan", 30, "999999999", "Data Innovate", Address{"Berlin", "", "Germany", "10117"}},
		{"Harper", 24, "111111111", "Global Designs", Address{"Sydney", "NSW", "Australia", "2000"}},
		{"Ryan", 26, "444444444", "Tech Experts", Address{"Toronto", "ON", "Canada", "999999"}},
		{"Evelyn", 25, "222222222", "Future Tech", Address{"Tokyo", "", "Japan", "10001"}},
		{"Jackson", 33, "888888888", "Smart Innovations", Address{"Dubai", "", "UAE", "12345"}},
		{"Chloe", 28, "777777777", "Innovative Systems", Address{"Paris", "", "France", "75001"}},
		{"Jack", 27, "333333333", "Tech Solutions Inc.", Address{"Moscow", "", "Russia", "10100"}},
		{"Lily", 29, "555555555", "Digital Ventures", Address{"Singapore", "", "Singapore", "49178"}},
		{"Owen", 32, "999999999", "Global Technologies", Address{"Seoul", "", "South Korea", "100011"}},
		{"Avery", 31, "666666666", "Digital Creations", Address{"Los Angeles", "CA", "USA", "90001"}},
		{"Zoe", 23, "222222222", "Tech Innovations", Address{"London", "", "UK", "999999"}},
		{"Carter", 24, "111111111", "Data Solutions", Address{"Berlin", "", "Germany", "10117"}},
		{"Madison", 32, "888888888", "Creative Tech", Address{"Sydney", "NSW", "Australia", "2000"}},
		{"Grayson", 26, "444444444", "Smart Designs", Address{"Toronto", "ON", "Canada", "999999"}},
		{"Hannah", 27, "333333333", "Innovative Solutions", Address{"Tokyo", "", "Japan", "10001"}},
		{"Wyatt", 28, "777777777", "Tech Experts Inc.", Address{"Dubai", "", "UAE", "12345"}},
		{"Ella", 30, "999999999", "Global Innovations", Address{"Paris", "", "France", "75001"}},
		{"Luke", 29, "555555555", "Digital Systems", Address{"Moscow", "", "Russia", "10100"}},
		{"Gabriella", 31, "666666666", "Creative Ventures", Address{"Singapore", "", "Singapore", "49178"}},
		{"Nathan", 25, "222222222", "Tech Co.", Address{"Seoul", "", "South Korea", "10011"}},
		{"Addison", 23, "222222222", "Data Tech", Address{"Los Angeles", "CA", "USA", "90001"}},
		{"Aria", 24, "111111111", "Smart Solutions", Address{"London", "", "UK", "333333"}},
		{"Julian", 32, "888888888", "Innovate Solutions", Address{"Berlin", "", "Germany", "10117"}},
		{"Sofia", 26, "444444444", "Global Tech Solutions", Address{"Sydney", "NSW", "Australia", "2000"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)

	allusers := []User{}
	for _, f := range records {
		employeeFound := User{}
		// if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
		// 	fmt.Println("Error", err)
		// }
		if err := bson.Unmarshal(f, &employeeFound); err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		allusers = append(allusers, employeeFound)
	}
	fmt.Println((allusers))

	// if err := db.Delete("users", "John"); err != nil {
	// 	fmt.Println("Error", err)
	// }

	// if err := db.Delete("users", ""); err != nil {
	// 	fmt.Println("Error", err)
	// }

}
