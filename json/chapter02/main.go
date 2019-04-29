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

func UnmarshalStruct() {
	str := "{\"name\":\"牛魔王\",\"age\":500,\"sal\":2000,\"skill\":\"牛魔拳\"}"

	var monster  Monster
	err := json.Unmarshal([]byte(str),&monster)
	if err != nil {
		fmt.Println("反序列化错误", err)
	}
	fmt.Println(monster)
	
}

func UnmarshalMap() {
	str := "{\"Age\":20,\"Name\":\"tony\",\"Sal\":2500,\"Skill\":\"JAVA\"}"
	var a map[string]interface{}
	err := json.Unmarshal([]byte(str),&a)
	if err != nil {
		fmt.Println("反序列化错误", err)
	}
	fmt.Println(a)
}


func UnmarshalSlice() {
	str := "[{\"Age\":20,\"Name\":\"tony\",\"Sal\":2500,\"Skill\":\"JAVA\"}," +
	"{\"Age\":18,\"Name\":\"tom\",\"Sal\":3000,\"Skill\":\"PHP\"}]"
	var a []map[string]interface{}
	err := json.Unmarshal([]byte(str),&a)
	if err != nil {
		fmt.Println("反序列化错误", err)
	}
	fmt.Println(a)
}

func main() {
	UnmarshalStruct()
	UnmarshalMap()
	UnmarshalSlice()
}