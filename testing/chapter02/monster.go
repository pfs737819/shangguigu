package monster
import (
	"fmt"
	"encoding/json"
	"os"
	"io/ioutil"
)
type Monster struct {
	Name string
	Age int
	Skill string
}

func (this *Monster) Store() string {
	

	data, err := json.Marshal(this)

	if err != nil {
		return fmt.Sprintf("Monster 序列化错误 %v", err) 
	}

	writeFilePath := "E:/goproject/src/go_code/shangguigu/testing/chapter02/monster.txt"

	Openfile, OpenFileErr := os.OpenFile(writeFilePath,os.O_WRONLY|os.O_CREATE,0666)

	defer Openfile.Close()

	if OpenFileErr != nil {
		return fmt.Sprintf("Monster os.OpenFile错误 %v", OpenFileErr) 
	}

	writeError := ioutil.WriteFile(writeFilePath,[]byte(data), 0666) 

	if writeError != nil {
		return fmt.Sprintf("Monster ioutil.WriteFile错误 %v", writeError) 
	}	
	
	return "nil"
}

func (this *Monster) ReStore() string {
	readFilePath := "E:/goproject/src/go_code/shangguigu/testing/chapter02/monster.txt"

	data, readeError := ioutil.ReadFile(readFilePath)
	if readeError != nil {
		return fmt.Sprintf("ioutil.ReadFile 错误%v", readeError)
	}


	UnmarshalError := json.Unmarshal([]byte(data),this)
	if UnmarshalError != nil {
		return fmt.Sprintf("json.Unmarshal 错误%v", UnmarshalError)
	}

	return "nil"
}

// func main(){
// 	monster := &Monster{
// 		Name:"牛魔王",
// 		Age:500,
// 		Skill:"牛角",
// 	}
// 	res := monster.Store()
// 	fmt.Println(res)
// }