package main

import "fmt"

// 人员档案
type Profile struct {
	id      int
	Name    string // 名字
	Age     int    // 年龄
	Married bool   // 已婚
}

//并且准备好了一堆原始数据，需要算法实现构建索引和查询的过程，代码如下：
func main() {
	list := []*Profile{
		{id: 1, Name: "张三", Age: 30, Married: true},
		{id: 2, Name: "李四", Age: 21},
		{id: 3, Name: "王麻子", Age: 21},
		{id: 4, Name: "张三1", Age: 30, Married: true},
		{id: 5, Name: "李四1", Age: 21},
		{id: 6, Name: "王麻子1", Age: 21},
		{id: 7, Name: "张三2", Age: 30, Married: true},
		{id: 8, Name: "李四2", Age: 21},
		{id: 9, Name: "王麻子2", Age: 21},
	}
	buildIndex2(list) //构建索引
	queryData2(5)     //查找
}

// 查询键
type queryKey struct {
	Name string
}

type queryKey2 struct {
	id int
}

// 创建查询键到数据的映射
var mapper = make(map[queryKey]*Profile)
var mapper2 = make(map[queryKey2]*Profile)

// 构建查询索引
func buildIndex2(list []*Profile) {
	// 遍历所有数据
	for _, profile := range list {
		// 构建查询键
		key := queryKey2{
			id: profile.id,
		}

		// 保存查询键
		mapper2[key] = profile
	}
	fmt.Println(mapper2)
}

func queryData2(id int) {
	// 根据查询条件构建查询键
	key := queryKey2{id}
	// 根据键值查询数据
	result, ok := mapper2[key]
	// 找到数据打印出来
	if ok {
		fmt.Println(result)
		fmt.Printf("%p", result)
	} else {
		fmt.Println("no found")
	}
}

//func buildIndex(list []*Profile) {
//	// 遍历所有数据
//	for _, profile := range list {
//		// 构建查询键
//		key := queryKey{
//			Name: profile.Name,
//		}
//
//		// 保存查询键
//		mapper[key] = profile
//	}
//	fmt.Println(mapper)
//}

// 根据条件查询数据
//func queryData(name string) {
//	// 根据查询条件构建查询键
//	key := queryKey{name}
//	// 根据键值查询数据
//	result, ok := mapper[key]
//	// 找到数据打印出来
//	if ok {
//		fmt.Println(result)
//		fmt.Printf("%p", result)
//	} else {
//		fmt.Println("no found")
//	}
//}
