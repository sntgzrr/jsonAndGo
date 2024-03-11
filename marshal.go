package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	Name     string
	Lastname string
	Age      int
}

func main() {
	/*----------------------Marshal----------------------*/
	p := person{
		Name:     "Santiago",
		Lastname: "Lozanette",
		Age:      21,
	}

	sb1, err1 := json.Marshal(p)
	if err1 != nil {
		fmt.Println("error", err1)
	}
	os.Stdout.Write(sb1)

	/*----------------------Unmarshal----------------------*/
	sb2 := []byte(`[
	{"Name": "Santiago", "Lastname": "Lozanette", "Age": 21},
	{"Name": "Rufus", 	 "Lastname": "Lozanette", "Age": 3}
]`)
	var persons []person
	err2 := json.Unmarshal(sb2, &persons)
	if err2 != nil {
		fmt.Println("error", err2)
	}
	fmt.Printf("%+v", persons)
}
