package main 

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct{
	Name string
	Surname string
	Tel []Telephone
}

type Telephone struct{
	Mobile bool
	Number string
}

func loadFromJson(filename string, key interface{}) error{
	in, err := os.Open(filename)
	if err!= nil{
		return err
	}

	decodeJson := json.NewDecoder(in)
	err = decodeJson.Decode(key)
	if err != nil{
		return err
	}
	in.Close()
	return nil
}

func saveToJson(filename *os.File, key interface{}){
	encodeJSON := json.NewEncoder(filename)
	err := encodeJSON.Encode(key)
	if err != nil{
		fmt.Println(err)
		return
	}
}

func main(){
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}

	filename := arguments[1]

	var myRecord Record
	err := loadFromJson(filename, &myRecord)
	if err == nil {
		fmt.Println(myRecord)
	}else{
		fmt.Println(err)
	}
}