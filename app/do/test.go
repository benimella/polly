package do

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"grpc/pb_file/order"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func F1() {
	a := &order.Order{
		Id:          100,
		Items:       nil,
		Description: "xxxx111",
		Price:       0,
		Destination: "xxxx555",
	}
	b, err := proto.Marshal(a)
	if err != nil {
		log.Fatalln(err)
	}
	filePath := GetCurrentDirectoryV2() + string(os.PathSeparator) + "tmp" + string(os.PathSeparator) + "1.txt"
	err = ioutil.WriteFile(filePath, b, 0666)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("写入成功")
	}
}

func F2()  {
	a := new(order.Order)
	filePath := GetCurrentDirectoryV2() + string(os.PathSeparator) + "tmp" + string(os.PathSeparator) + "1.txt"
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	err = proto.Unmarshal(b, a)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", a)
}

func GetCurrentDirectoryV1() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func GetCurrentDirectoryV2() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	path = path[:index]
	return path
}
