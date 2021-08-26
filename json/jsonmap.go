package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b := []byte(`{"IP": "127.0.0.1", "name": "SKY"}`)

	m := make(map[string]string)

	err := json.Unmarshal(b, &m)
	if err != nil {

		fmt.Println("Umarshal failed:", err)
		return
	}

	fmt.Println("m:", m)

	for k,v :=range m {
		fmt.Println(k, ":", v)
	}
}
