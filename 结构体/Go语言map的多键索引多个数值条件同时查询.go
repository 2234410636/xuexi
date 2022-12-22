package main

import "fmt"

// 人员档案
type Profile struct {
	Name    string // 名字
	Age     int    // 年龄
	Married bool   // 已婚
}

//并且准备好了一堆原始数据，需要算法实现构建索引和查询的过程，代码如下：
func main() {
	list := []*Profile{
		{Name: "张三", Age: 30, Married: true},
		{Name: "李四", Age: 21},
		{Name: "王麻子", Age: 21},
		{Name: "张三1", Age: 30, Married: true},
		{Name: "李四1", Age: 21},
		{Name: "王麻子1", Age: 21},
		{Name: "张三2", Age: 30, Married: true},
		{Name: "李四2", Age: 21},
		{Name: "王麻子2", Age: 21},
	}
	buildIndex(list) //构建索引
	queryData("张三2") //查找
}

// 查询键
type queryKey struct {
	Name string
}

// 创建查询键到数据的映射
var mapper = make(map[queryKey]*Profile)

// 构建查询索引
func buildIndex(list []*Profile) {
	// 遍历所有数据
	for _, profile := range list {
		// 构建查询键
		key := queryKey{
			Name: profile.Name,
		}

		// 保存查询键
		mapper[key] = profile
	}
	fmt.Println(mapper)
}

// 根据条件查询数据
func queryData(name string) {
	// 根据查询条件构建查询键
	key := queryKey{name}
	// 根据键值查询数据
	result, ok := mapper[key]
	// 找到数据打印出来
	if ok {
		fmt.Println(result)
		fmt.Printf("%p", result)
	} else {
		fmt.Println("no found")
	}
}
