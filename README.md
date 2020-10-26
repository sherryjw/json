# 对象序列化支持包
## 概述
&emsp;&emsp;将一个对象写成特定文本格式的字符流，即序列化。序列化通过某种存储形式使自定义对象持久化，使对象在不同平台、应用程序等的传递更加方便简洁。<br/>
&emsp;&emsp;本程序包 ``json`` 提供 ``JsonMarshal`` 函数将结构数据格式化为 json 字符流。<br/>
<br/>

## 安装

```R
go get github.com/sherryjw/json
```
<br/>

## 使用说明
&emsp;&emsp;下面通过一个简单的例子来介绍如何使用该程序包提供的函数。<br/>

&emsp;&emsp;在工作目录/main下创建 main.go，编辑代码如下：<br/>
```go
package main

import (
	"fmt"
	"os"
	json "github.com/sherryjw/json"
)

func main() {	
	type ColorGroup struct {
		Num         []int        `json:"num"`
		Right       bool         `json:"-,"`
		ID          int          `json:"id,omitempty"`
		Name        string       `json:"-"`
		Colors      []string     `json:",omitempty"`
		Dictionary  map[int]int  `json:"dictionary,omitempty"`
	}

	group := ColorGroup{
		Num:	    []int{55, 8, -12},
		Name:       "Reds",
		Right:	    true,
		Colors:     []string{"Crimson", "<15", "Ruby"},
		Dictionary: map[int]int{16: 1, 75:100},
	}

	b, err := Json.JsonMarshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
```

&emsp;&emsp;运行：
```R
go run main.go
```

&emsp;&emsp;结果如下：
```R
{"num":[55,8,-12],"-":true,"Colors":["Crimson","\u003c15","Ruby"],"dictionary":{"16":1,"75":100}}
```

&emsp;&emsp;正如我们所期待的那样，结构体类型的变量 ``group`` 的数据被序列化为一串 json 字符流。

<br/>

## API 文档

[![GoDoc](https://img.shields.io/badge/GoDoc-Reference-blue?style=for-the-badge&logo=go)]([doc_zh_CN.md](https://sherryjw.gitee.io/json/json%20-%20Go%20Documentation%20Server.html))
