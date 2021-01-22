package main

import (
	"fmt"
	"log"

	"github.com/kulwindermatharu/excelize"
)

//M ..
type M map[string]interface{}

var data = []M{
	{"Name": "1Noval", "Gender": "male", "Age": 18},
	{"Name": "1Nabila", "Gender": "female", "Age": 12},
	{"Name": "1Yasa", "Gender": "male", "Age": 11},
	{"Name": "1Noval", "Gender": "male", "Age": 18},
	{"Name": "1Nabila", "Gender": "female", "Age": 12},
	{"Name": "1Yasa", "Gender": "male", "Age": 11},
}

func main() {
	xlsx := excelize.NewFile()
	sheet1Name := "Sheet One"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)
	xlsx.SetCellValue(sheet1Name, "A1", "Name")
	xlsx.SetCellValue(sheet1Name, "B1", "Gender")
	xlsx.SetCellValue(sheet1Name, "C1", "Age")
	err := xlsx.AutoFilter(sheet1Name, "A1", "C1", "")
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}
	for i, each := range data {
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), each["Name"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), each["Gender"])
		xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), each["Age"])
	}
	err = xlsx.SaveAs("file.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
