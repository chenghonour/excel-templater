package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"

	"github.com/chenghonour/excel-templater"
)

var (
	templateFile = "./examples/template.xlsx"
	resultFile   = "./examples/result.xlsx"
	useDefault   = false

	//go:embed payload.json
	data []byte
)

func main() {
	templater := excel.NewTemplater(useDefault)

	var payload interface{}
	if err := json.Unmarshal(data, &payload); err != nil {
		panic(err)
	}

	fileByte, err := templater.FillIn(templateFile, payload)
	if err != nil {
		panic(err)
	}

	// save bytes to file
	err = os.WriteFile(resultFile, fileByte, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}
