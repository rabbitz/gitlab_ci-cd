package main

import (
	"fmt"

	// import里引入的根目录名称是工程名称，也就是go.mod里的module名称
	controllers "github.com/rabbitz/basis/handlers"

	jsoniter "github.com/json-iterator/go"
)

type AppInfo struct {
	Name string
}

func main() {
	info := AppInfo{
		Name: "GoApp",
	}
	jsonString, _ := jsoniter.Marshal(&info)

	fmt.Println(string(jsonString))
	controllers.ShowInfo()
}
