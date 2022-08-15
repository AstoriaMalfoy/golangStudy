package main

import (
	"encoding/json"
	"fmt"
)

type Persion struct {
	Name    string  `json:"name,omitempty"`
	Age     int64   `json:"age"`
	Weight  float64 `json:"-"`
	Profile Profile
}

type Profile struct {
	Website string `json:"website,omitempty"`
	Slogan  string `json:"slogan,omitempty"`
}

func main() {
	qSerializerDemo()
	fmt.Println("/////////////////////")
	serializerDemo()
}

func qSerializerDemo() {
	prof := Profile{
		Website: "http://testdemo.com",
		Slogan:  "this is test profile",
	}

	per := Persion{
		Name:    "astoria",
		Age:     24,
		Weight:  12.2,
		Profile: prof,
	}

	jsonByte, err := json.Marshal(per)
	if err != nil {
		fmt.Printf("json marshal error err %v\n", err)
		return
	}
	var p2 Persion
	err = json.Unmarshal(jsonByte, &p2)
	if err != nil {
		fmt.Printf("json unMarshal error err %v\n", err)
		return
	}
	fmt.Printf("the json marshal p2:%#v\n", p2)
}

func serializerDemo() {
	p1 := Persion{
		Name:   "astoria",
		Age:    12,
		Weight: 12.2,
	}
	byte, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal failed, error %v\n", err)
	}
	fmt.Printf("str:%s\n", byte)

	var p2 Persion
	err = json.Unmarshal(byte, &p2)
	if err != nil {
		fmt.Printf("json.Unmarshal failed , err %v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}
