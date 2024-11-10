package utils

import (
	"encoding/json"
	"fmt"
)

func PrintJson(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Printf("%s\n", jsonData)
}
