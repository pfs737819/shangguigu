
package main
import (
	"fmt"
)

//二分查找的函数
/*
二分查找的思路: 比如我们要查找的数是 findVal
1. arr是一个有序数组，并且是从小到大排序
2.  先找到 中间的下标 middle = (leftIndex + rightIndex) / 2, 然后让 中间下标的值和findVal进行比较
2.1 如果 arr[middle] > findVal ,  就应该向  leftIndex ---- (middle - 1)
2.2 如果 arr[middle] < findVal ,  就应该向  middel+1---- rightIndex
2.3 如果 arr[middle] == findVal ， 就找到
2.4 上面的2.1 2.2 2.3 的逻辑会递归执行
3. 想一下，怎么样的情况下，就说明找不到[分析出退出递归的条件!!]
if  leftIndex > rightIndex {
   // 找不到..
   return ..
}
*/

// func findNum(arr *[8]int,leftIndex int,rightIndex int,findVal int){
// 	fmt.Println("第一步",leftIndex,rightIndex)

// 	if leftIndex > rightIndex {
// 		fmt.Println("找不到",leftIndex,rightIndex)
// 		return
// 	}
// 	middelIndex := (leftIndex + rightIndex) / 2
// 	fmt.Println("middelIndex",middelIndex)
// 	if arr[middelIndex] > findVal {
// 		findNum(arr,leftIndex,middelIndex - 1,findVal)
// 	}else if arr[middelIndex] < findVal {
// 		findNum(arr,middelIndex + 1,rightIndex,findVal)
// 	}else{
// 		fmt.Printf("找到数字%v的位置为%v",findVal,middelIndex)
// 	}

// }

type Point struct{
	X int
	Y int
} 
func main() {

	// var arr [4][4]int
	// arr[0][0] = 1
	// arr[1][1] = 2
	// arr[2][2] = 3
	// arr[3][3] = 4

	// arr[0][3] = 4
	// arr[1][2] = 3
	// arr[2][1] = 2
	// arr[3][0] = 1

	// for i := 0;i < 4;i++ {
	// 	for j := 0;j < 4; j++ {
	// 		fmt.Print(arr[i][j]," ")
	// 	}
	// 	fmt.Println()
	// }

	// student := make(map[string]map[string]string)
	// student["01"] = make(map[string]string,3)
	// student["01"]["name"] = "tom"
	// student["01"]["age"] = "21"
	// student["01"]["sex"] = "男"

	// student["02"] = make(map[string]string,3)

	// student["02"]["name"] = "jack"
	// student["02"]["age"] = "21"
	// student["02"]["sex"] = "女"

	// type stu struct{
	// 	name string
	// 	age int
	// 	address string
	// }

	// student := make(map[string]stu)
	// student["no1"] = stu{"tom",18,"上海"}
	// student["no2"] = stu{"jack",20,"北京"}
	
	// for i,v := range student{
	// 	fmt.Printf("编号:%v,姓名:%v,年龄:%v,地址:%v\n",i,v.name,v.age,v.address)
	// }
	// var catName string

	// cat := make(map[string]map[string]string)
	// cat["小白"] = make(map[string]string,3)
	// cat["小白"]["年龄"] = "3"
	// cat["小白"]["颜色"] = "白色"

	// cat["小花"] = make(map[string]string,3)
	// cat["小花"]["年龄"] = "100"
	// cat["小花"]["颜色"] = "花色" 

	// fmt.Println("请输入小猫的名字")
	// fmt.Scanln(&catName)

	// if cat[catName] != nil {
	// 	fmt.Printf("小猫的名字为%v,年龄为%v,颜色为%v\n",catName,cat[catName]["年龄"],cat[catName]["颜色"])	
	// }

	// fmt.Println(cat)
	
	// for i,v := range student {
	// 	fmt.Println(i,v)
	// 	for i1,v1 := range student[i] {
	// 		fmt.Println(i1,v1)
	// 	}
		
	// }



	// var r1 Point
	var r1 Point = Point{11,22}

	var r2 Point = Point{13,23}
	// fmt.Printf("r1地址%p\n",&r1)
	fmt.Printf("r1X地址 %p\n",&r1.X)
	fmt.Printf("r2Y地址 %p\n",&r2.Y)
	
	

}