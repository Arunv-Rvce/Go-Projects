package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Employee struct {
	id              string `json:"id"`
	employee_name   string `json:"employee_name"`
	employee_salary string `json:"employee_salary"`
	employee_age    string `json:"employee_age"`
	profile_image   string `json:"profile_image"`
}

var client = &http.Client{}

func get_example() {
	resp, _ := client.Get("http://dummy.restapiexample.com/api/v1/employees")
	defer resp.Body.Close()
	var decoded []Employee
	content, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(content, &decoded)
	if err == nil {
		for i := 0; i < len(decoded); i++ {
			fmt.Println(decoded[i])
		}
	} else {
		fmt.Println(err)
	}
}
