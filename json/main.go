package main
import (
	"fmt"
)

type Monster struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func (monster *Monster) String()string{
	monster.Name = "青牛精"
	return "110"
}

func main() {
	var monster Monster
	fmt.Println(&monster)
	// monster := monster{"牛魔王",500,"牛角"}
	// strJson , err := json.Marshal(monster)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(monster,string(strJson))
}