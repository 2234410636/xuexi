package main

import "fmt"

//func main() {
//	a := make(map[string]map[string]map[string]string)
//	a["1"] = make(map[string]map[string]string)
//	a["1"]["1"] = make(map[string]string)
//	a["1"]["2"] = make(map[string]string)
//	a["1"]["1"]["1"] = "tom"
//	a["1"]["2"]["1"] = "tm"
//	fmt.Println(a["1"]["1"]["1"])
//	fmt.Println(a)
//}
//func DeleteMap(m map[int]string, key int) {
//	delete(m, key) //删除key值为2的map
//	for k, v := range m {
//		fmt.Printf("len(m)=%d, %d ----> %s\n", len(m), k, v)
//
//	}
//	//len(m)=2, 1 ----> Luffy
//	//len(m)=2, 3 ----> Zoro
//}
//
//func main() {
//	m := map[int]string{1: "Luffy", 2: "Sanji", 3: "Zoro"}
//	DeleteMap(m, 2) //删除key值为2的map
//
//	for k, v := range m {
//		fmt.Printf("len(m)=%d, %d ----> %s\n", len(m), k, v)
//
//	}
//	//len(m)=2, 1 ----> Luffy
//	//len(m)=2, 3 ----> Zoro
//}

//func test() map[int]string {
//	// m1 := map[int]string{1: "Luffy", 2: "Sanji", 3: "Zoro"}
//	m1 := make(map[int]string, 1) // 创建一个初创容量为1的map
//	m1[1] = "buffy"
//	m1[2] = "snji" // 自动扩容
//	m1[67] = "oro"
//	m1[2] = "ami" // 覆盖 key值为2 的map
//	fmt.Println("m1 = ", m1)
//	return m1
//}
//
//func main() {
//	m2 := test() // 返回值 —— 地址传递
//	fmt.Println("m2 = ", m2)
//	m3 := make([]string, 0)
//	for _, v := range m2 {
//		m3 = append(m3, v)
//	}
//	sort.Strings(m3)
//	fmt.Println("m3 = ", m3)
//}

func main() {
	users := make(map[string]map[string]string)
	users["smith"] = make(map[string]string)
	users["smith"]["password"] = "999999"
	users["smith"]["nickname"] = "小花猫"

	var name string
	fmt.Scanf("%s\n", &name)
	modifyUser(users, name)
	fmt.Println(users)

}
func modifyUser(users map[string]map[string]string, name string) {
	if users[name] != nil {

		users[name]["password"] = "88888888"
	} else {
		users[name] = make(map[string]string)
		var nickname, password string
		fmt.Scanf("%s\n", &nickname)
		users[name]["nickname"] = nickname
		fmt.Scanf("%s\n", &password)
		users[name]["password"] = password
	}

}
