package main

import "fmt"

type Person struct {
	Name string
	Age  int
	City string
	//email 추가
	Email string
}

func (p Person) Introduce() string {
	return fmt.Sprintf("저는 %s이고 %d살입니다. email은 %s입니다. 사는 곳은 %s입니다.", p.Name, p.Age, p.Email, p.City)
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func (p *Person) HaveBirthday() {
	p.Age++
}

func main() {
	person := Person{
		Name: "Alice",
		Age:  20,
		City: "Seoul",
		Email: "alice@example.com",
	}

	fmt.Println(person.Introduce())

	if person.IsAdult() {
		fmt.Println(person.Name, "is an adult.")
	}

	person.HaveBirthday()
	fmt.Println("after birthday:", person.Age)
}
