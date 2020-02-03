package main

import "fmt"

type USBDevice interface {
	Connect() bool
	ReadData() []byte
	Disconnect() bool
}

type Upan struct {
	//定义容量
	storage int
}

func (up *Upan) Connect() bool {
	fmt.Println("优盘连接成功！")
	return true
}

func (up *Upan) ReadData() []byte {
	fmt.Println("读取数据！")
	return []byte("无码.avi")
}

func (up *Upan) Disconnect() bool {
	fmt.Println("优盘断开成功！")
	return true
}

func main() {

	var usbdevice USBDevice

	var up = Upan{8 * 1024 * 1024 * 1024}

	usbdevice = &up

	usbdevice.Connect()
	fmt.Println(usbdevice.ReadData())
	usbdevice.Disconnect()

}
