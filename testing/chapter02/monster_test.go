package monster
import (
	// "fmt"
	"testing"
)

// func TestStore(t *testing.T) {
// 	monster := &Monster{
// 		Name:"牛魔王",
// 		Age:500,
// 		Skill:"牛角",
// 	}
// 	res := monster.Store()
// 	if res != "nil" {
// 		// fmt.Printf("AddUpper(10) 执行错误, 期望值%v 实际值%v\n", 55, res)
// 		t.Fatalf("%v\n", "", res)
// 	}else{
// 		t.Logf("测试通过")
// 	}
// }

func TestReStore(t *testing.T) {
	monster := &Monster{}
	res := monster.ReStore()
	if res != "nil" {
		// fmt.Printf("AddUpper(10) 执行错误, 期望值%v 实际值%v\n", 55, res)
		t.Fatalf("%v\n", "", res)
	}else{
		t.Logf("测试通过%v",*monster)
	}
}
