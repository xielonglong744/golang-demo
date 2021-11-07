package main

import (
	"fmt"
	"github.com/xielonglong744/go_micro/protobuf_demo/address"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
)

func main() {
	var cb address.ContackBook
	p1 := address.Person{
		Name: "xielonglong",
		Gender: address.GenderType_MALE,
		Number: "779777",
	}
	fmt.Println(p1.Name)
	cb.Persons = append(cb.Persons,&p1)

	data, err := proto.Marshal(&p1)
	if err != nil {
		fmt.Println("proto marshal failed")
		return
	}
	err = ioutil.WriteFile("./proto.data",data,0644)
	if err != nil {
		fmt.Println("writefile failed ")
		return
	}
	data2, err := ioutil.ReadFile("./proto.data",)
	if err != nil {
		fmt.Println("readfile failed ")
		return
	}
	var p2 address.Person
	proto.Unmarshal(data2,&p2)
	fmt.Println(p2.Name)
}

