package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	inputData := `[{"key":"key1","value":1},{"key":"key2","value":2},{"key":"key1","value":3}]`

	var items []map[string]interface{}

	if err := json.Unmarshal([]byte(inputData), &items); err != nil {
		log.Fatal(err)
	}

	processedData := make(map[string]float64)
	for _, item := range items {
		key := item["key"].(string)
		value := item["value"].(float64)

		processedData[key] = value
	}

	outputData, err := json.Marshal(processedData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(outputData))
}
