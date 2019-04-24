package main
import (
	"fmt"
)

type Usb interface{
	Start()
	Stop()
}

type Phone struct{

}

func(p Phone) Start(){
	fmt.Println("Phone开始运行")
}

func(p Phone) Call(){
	fmt.Println("Phone开始打电话")
}

func(p Phone) Stop(){
	fmt.Println("Phone停止运行\n")
}



type Ipad struct{

}

func(i Ipad) Start(){
	fmt.Println("Ipad开始运行")
}

func(i Ipad) Stop(){
	fmt.Println("Ipad停止运行\n")
}



type Computer struct{

}

func(c Computer) Working(usb Usb){
	usb.Start()
	usb.Stop()
	usb.Call()
}

func main() {

	Phone := Phone{}
	Ipad := Ipad{}
	Computer := Computer{}
	Computer.Working(Ipad)
	Computer.Working(Phone)
}