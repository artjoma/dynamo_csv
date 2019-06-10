package main
/**
	AWS DynamoDB universal json to csv converter
	Artjom Aminov
	07.06.2019
 */
import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Model struct {
	Count uint32 								`json:"Count"`
	Items []map[string]map[string]interface{}	`json:"Items"`
}

func main (){
	if len(os.Args) != 3 {
		panic("first arg source file name, second to file name")
	}
	parse(os.Args[1], os.Args[2])
}

func parse(sourceFileName, toFileName string){
	fmt.Println(time.Now(), "start")
	fmt.Println(sourceFileName, "->", toFileName)

	//read to JSON model
	pwd, _ := os.Getwd()
	fmt.Println("pwd:", pwd)
	fromFile, err := ioutil.ReadFile(pwd + "/" + sourceFileName)
	if err != nil{
		panic("can't read from source file")
	}
	model := &Model{}
	err = json.Unmarshal(fromFile, model)
	if err != nil {
		panic("err parse source file. err:" + err.Error())
	}

	fmt.Println("count:", model.Count, "/", len(model.Items))

	//make header
	csvData := make([][]string, model.Count + 1)
	var headerRow []string = nil
	for key, _ := range model.Items[0]  {
		//add csv header
		fmt.Println("col:", key)
		headerRow = append(headerRow, key)
	}
	csvData[0] = headerRow

	//make body
	i := 1
	var csvRow []string = nil
	for _, value := range model.Items {
		csvRow = make([]string, 0, 20)
		for _, header := range headerRow {
			for _, value := range value[header] {
				if value == nil{
					csvRow = append(csvRow, "")
				}else{
					csvRow = append(csvRow, fmt.Sprint(value))
				}
			}
		}
		csvData[i] = csvRow
		i++
	}

	//save csv file
	osToFile, err := os.Create(pwd + "/" + toFileName)
	writer := csv.NewWriter(osToFile)
	writer.Comma = ';'
	writer.WriteAll(csvData)
	writer.Flush()
	osToFile.Close()

	fmt.Println(time.Now(), "end")
}
