package cal
import (
	// "fmt"
	"testing"
)
func TestGetSum(t *testing.T) {
	res := GetSum(50,5)
	if res != 55 {
		// fmt.Printf("AddUpper(10) 执行错误, 期望值%v 实际值%v\n", 55, res)
		t.Fatalf("GetSum(50,10) 执行错误, 期望值%v 实际值%v\n", 55, res)
	}else{
		t.Logf("测试通过")
	}
}