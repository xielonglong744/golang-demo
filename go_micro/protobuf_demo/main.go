package main

import (
	"fmt"
	"github.com/xielonglong744/go_micro/protobuf/address"
)

func main() {
	var cb address.ContactBook
	p1 := address.Person{
		Name: "xielonglong",
		Gender: address.GenderType_MALE,
		Number: "779777",
	}
	fmt.Println(p1)


}

