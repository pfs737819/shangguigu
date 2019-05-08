package test
import (
	"testing"
	"reflect"
	"fmt"
)

type User struct {
	UserID string
	Name string
}

func TestResultFunc(t *testing.T) {
	var (
		model *User
		ftype reflect.Type
		elem  reflect.Value
	)

	ftype = reflect.TypeOf(model)
	fmt.Printf("ftype %v\n",ftype.Kind())//指针类型
	ftype = ftype.Elem()
	fmt.Printf("ftype %v\n",ftype.Kind())//struct类型

	elem = reflect.New(ftype)
	fmt.Printf("ftype %v\n",elem.Kind().String())
	fmt.Printf("ftype %v\n",elem.Elem().Kind().String())

	model = elem.Interface().(*User) //model是*user 它的指向和elem是一样的
	elem = elem.Elem()
	elem.FieldByName("UserID").SetString("123456")
	elem.FieldByName("Name").SetString("tom")
	fmt.Printf("elem %v UserID %v Name %v\n",model,model.UserID,model.Name)

	
}