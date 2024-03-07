/*
[
	{"name": "David", "age": 28},
	{"name": "Eva", "age": 35},
	{"name": "Frank", "age": 22},
	{"name": "Grace", "age": 27},
	{"name": "Henry", "age": 31},
	{"name": "Ivy", "age": 23},
	{"name": "Jack", "age": 29},
	{"name": "Kelly", "age": 26},
	{"name": "Liam", "age": 33},
	{"name": "Mia", "age": 24}
]

Suppose we have a 'data.json' like this, and after executing the code, we get the below response:

People under 30:
Name: Frank, Age: 22
Name: Ivy, Age: 23
Name: Mia, Age: 24
Name: Kelly, Age: 26
Name: Grace, Age: 27
Name: David, Age: 28
Name: Jack, Age: 29

*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Read JSON data from file
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal JSON data into a slice of Person structs
	var people []Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Sort people by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	// Filter people under 30 years old
	var youngPeople []Person
	for _, person := range people {
		if person.Age < 30 {
			youngPeople = append(youngPeople, person)
		}
	}

	// Print the filtered and sorted list of people
	fmt.Println("People under 30:")
	for _, person := range youngPeople {
		fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	}
}
