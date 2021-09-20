package main

import (
	"codeGen/structs"
	"fmt"
)

func main() {

	testJSON := `{"ID":123, "Name":"Ivan", "Login":"IvanDev", "Password":"1234567890", "Email":"ivan@gmail.com", "Status":"Admin"}`
	fmt.Println(testJSON)

	user := &structs.User{}
	user.UnmarshalJSON([]byte(testJSON))
	fmt.Printf("%+v\n", user)

	user.Status = "Editor"
	outJSON, _ := user.MarshalJSON()
	fmt.Println(string(outJSON))

}
