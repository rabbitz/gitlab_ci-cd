

#### 声明

`package 包名`，注意事项：

* 一个文件夹下直接包含的文件只能归属一个package，同样一个package不能在多个文件夹下
* 包名和文件夹名字无关，包名不能包含 `-` 符号
* 包名为 `main` 的包为应用程序的入口包，这种包边以后会得到一个可执行文件，而编译不包含 `main ` 包的源代码则不会得到可执行文件

#### 可见性

引用包里的标识符（如变量，常量，类型，函数等）时，该标识符必须对外可见（public），只需要将标识符的 **`首字母大写 `**就可以让标识符对外可见。

* 函数局部变量首字母大写，外部不可见，只能当前函数使用

* 结构体中的字段名和接口中的方法名如果首字母大写，外部包可以访问这些字段和方法

  ```go
  type Student struct {
  	Name  string //可在包外访问的方法
  	class string //仅限包内访问的字段
  }
  
  type Payer interface {
  	init() //仅限包内访问的方法
  	Pay()  //可在包外访问的方法
  }
  ```

#### 包的导入

`import "包的路径"` ，注意事项：

* import 导入语句通常放在文件开头包声明语句的下面
* 导入的包名需要用双引号包裹
* 包名是从`$GOPATH/src/`后开始计算的，使用`/`进行路径分隔。
* Go语言中禁止循环导入包

#### 自定义包名

`import 别名 "包的路径"` ，别名不需要加引号

#### 匿名包导入

`import _ "包的路径"` 只导入包，不使用包内部数据。会被编译到可执行文件中

#### init()初始化函数

##### init()函数介绍

在Go语言程序执行时导入包语句会自动触发包内部`init()`函数的调用。需要注意的是： `init()`函数没有参数也没有返回值。 `init()`函数在程序运行时自动被调用执行，不能在代码中主动调用它。

包初始化执行的顺序如下图所示：![包中的init()执行时机](https://www.liwenzhou.com/images/Go/package/init01.png)

##### init()函数执行顺序

Go语言包会从`main`包开始检查其导入的所有包，每个包中又可能导入了其他的包。Go编译器由此构建出一个树状的包引用关系，再根据引用顺序决定编译顺序，依次编译这些包的代码。

在运行时，被最后导入的包会最先初始化并调用其`init()`函数， 如下图示：![包之间的init()执行顺序](https://www.liwenzhou.com/images/Go/package/init02.png)

#### 使用 Module 配置

##### 在同一项目下

目录结构如下：

```go
moduledemo
├── go.mod
├── main.go
└── mypackage
  └── mypackage.go
```

`moduledemo/go.mod` 中定义如下：

```go
module moduledemo

go 1.15
```

`mypackage.go` 中包引入方式为：

```go
// 模块名 + 引入的包名
import (
 "fmt"
 "moduledemo/mypackage" // 导入同一项目下的mypackage包
)
```

##### 不同项目下

目录结构如下：

```go
├── moduledemo
│   ├── go.mod
│   └── main.go
└── mypackage
    ├── go.mod
    └── mypackage.go
```

`moduledemo/go.mod` 中定义如下：

```go
module moduledemo

go 1.15

// 需要手动添加此行
replace mypackage => ../mypackage 

// 手动添加此行
// 如若没添加 go run/build 时会自动加入：
// require mypackage v0.0.0-00010101000000-000000000000
// 同样这里的 mypackage 根据实际 .mod 文件中定义做修改
require "mypackage" v0.0.0
```

```go
import (
	"fmt"
	"mypackage"
)
```

`mypaackage/go.mod` 中定义如下：

```go
module mypackage

go 1.15
```



