package main
import (
	"fmt"
	_"math/rand"
	_"time"
	// "go_code/project01/public"
)

// func peach(n int) int {
// 	if n == 10 {
// 		return 1
// 	}else{
// 		return (peach(n + 1) + 1) * 2 
// 	}
// }

// func sum(n1 int,vars... int) (res int) {
// 	sum := n1
// 	fmt.Println(sum)
// 	for i := 0; i < len(vars);i ++ {
// 		sum += vars[i]
// 	}
// 	res = sum
// 	return
// }

// func swap (n1 *int ,n2 *int){
// 	t := *n1
// 	*n1 = *n2
// 	*n2 = t
// 	fmt.Printf("n1=%v n2=%v\n",*n1,*n2)
// }

// func init() {
// 	fmt.Println("init()")
// }

/////////////////闭包///////////////////
// func AddSupper() func (int) int {
// 	n := 1
// 	return func (x int) int {
// 		n += x
// 		return n
// 	}
// }
////////////////////////////////////////

// func makeSuffix
// func ResDate() string {
// 	var year int
// 	var month int
// 	var day int
// 	fmt.Println("请输入年：")
// 	fmt.Scanln(&year)
// 	if(year > 0){

// 	}
// }

// func modify(arr[3] int){
// 	arr[0] = 100
// 	fmt.Println("modify的arr",arr)
// }

// func later(){
// 	var zimuArr [26] byte
// 	for i := 0 ; i < 26 ; i ++{
// 		zimuArr[i] = 'A'+byte(i)
// 		fmt.Printf("%c",zimuArr[i])
// 	}
// 	fmt.Println(zimuArr)
// }

// func ArrFlip(){
// 	var arr [5]int
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0;i < 5;i ++ {
// 		arr[i] = rand.Intn(100)
// 	}

// 	len := len(arr)
// 	var temp int

// 	fmt.Println(arr)

// 	for i := 0;i < len/2;i ++ {
// 		temp = arr[len-1-i]
// 		arr[len-1-i] = arr[i]
// 		arr[i] = temp
		
// 	}

// 	fmt.Println(arr)
// }

// func SliceArr(){
// 	var arr [5]int = [5]int{1,2,3,4,5}
// 	fmt.Println(arr)
// 	var SliceArr [] int = arr[:]
// 	var makeArr [] int = make([]int,1,2)
// 	makeArr = append(makeArr,SliceArr...)
// 	makeArr2 := makeArr
// 	makeArr2[0] = 100
// 	fmt.Println(makeArr)
// 	for i,v := range SliceArr{
// 		fmt.Printf("i=%v,v=%v\n",i,v)
// 	}
// 	fmt.Println(SliceArr)
// }

// func fbn(n int) []uint64 {
// 	SliceArr := make([]uint64,n)
// 	SliceArr[0] = 1
// 	SliceArr[1] = 1
// 	for i := 2;i < n;i ++ {
// 		SliceArr[i] = SliceArr[i-1] + SliceArr[i-2]
// 	}

// 	return SliceArr
// }
func bubbleSort(arr *[5]int){
	var temp int
	for j := 0;j < len(*arr) - 1;j ++ {
		if (*arr)[j] > (*arr)[j+1] {
			temp = (*arr)[j]
			(*arr)[j] = (*arr)[j+1]
			(*arr)[j+1] = temp
		}	
	}

	for j := 0;j < len(*arr) - 1;j ++ {
		if (*arr)[j] > (*arr)[j+1] {
			temp = (*arr)[j]
			(*arr)[j] = (*arr)[j+1]
			(*arr)[j+1] = temp
		}	
	}
}

func main(){ 
	// res1 := public.OutPututils();
	// fmt.Println(res1)
	// var sum int 
	// var price float64 = 2000
	// for {
	// 	if price > 50000 {
	// 		price = price-price*0.05
	// 	}else{
	// 		if price >= 1000 {
	// 			price -= 1000
	// 		}else{
	// 			fmt.Println(price)
	// 			break
	// 		}
	// 	}

	// 	sum ++
	// }

	// fres := func(n1 int , n2 int) int {
	// 	return n1 + n2
	// }

	// p := fres

	// var n1 int = 1
	// var n2 int = 2
	// swap(&n1,&n2)
	// fmt.Printf("n1=%v n2=%v\n",n1,n2)
	// fmt.Println(p(1,2))
	// res := AddSupper()
	// fmt.Println(res(1))
	// fmt.Println(res(1))
	// fmt.Println(res(1))
	// fmt.Println(res(1))
	// var arr = [...]int{1,2,3}
	// var str string = "Hello 中国"
	// str1 := []rune(str)
	// fmt.Printf("%T,%c",str[1],str1[7])
	// fbnSlice := fbn(10)
	var arr[5] int = [5]int{10,2,1,7,6}
	bubbleSort(&arr)
	fmt.Println(arr)
}