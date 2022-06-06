package main 

import (
	"fmt"
	 "encoding/json"
)

type Student struct {
	Name  string 
	Points float64
}

func main() {
	singleStudent:= `{"Name": "Munashe", "Points": 45.8}` 

	var final Student 
     
	err:= json.Unmarshal([]byte(singleStudent), &final)  

	if(err != nil) {
		fmt.Println(err)
	}
	fmt.Printf("%+v", final)
	
}
