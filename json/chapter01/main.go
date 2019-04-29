package main
import (
	"fmt"
	"encoding/json"
)

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Sal float64 `json:"sal"`
	Skill string `json:"skill"`
}

func testSlice() {
	var a []map[string]interface{}

	var m1 map[string]interface{} 
	m1 = make(map[string]interface{})
	m1["Name"] = "tony"
	m1["Age"] = 20
	m1["Sal"] = 2500.00
	m1["Skill"] = "JAVA"

	a = append(a,m1)

	var m2 map[string]interface{} 
	m2 = make(map[string]interface{})
	m2["Name"] = "tom"
	m2["Age"] = 18
	m2["Sal"] = 3000.00
	m2["Skill"] = "PHP"
	a = append(a,m2)
	jsonData, err := json.Marshal(a) 
	if err != nil {
		fmt.Println("序列化失败", err)
	}

	fmt.Println(string(jsonData))
}

func testMap() {
	var a map[string]interface{}

	a = make(map[string]interface{})
	a["Name"] = "tony"
	a["Age"] = 20
	a["Sal"] = 2500.00
	a["Skill"] = "JAVA"

	jsonData, err := json.Marshal(a) 
	if err != nil {
		fmt.Println("序列化失败", err)
	}

	fmt.Println(string(jsonData))
}

func main() {
	monster := Monster {
		Name:"牛魔王",
		Age:500,
		Sal:2000.00,
		Skill:"牛魔拳",
	}

	jsonData, err := json.Marshal(monster)

	if err != nil {
		fmt.Println("序列化失败", err)
	}
	fmt.Println(string(jsonData))

	testMap()
	testSlice()
}