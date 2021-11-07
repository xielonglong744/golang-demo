package main

//func main() {
//	a := "sssssssss"
//	fmt.Printf("%v\n",a)
//	fmt.Printf("%s\n",a)
//	fmt.Printf("%p\n",a)
//	fmt.Printf("%T\n",a)
//	fmt.Printf("%t\n",a)
//}
//
//func main() {
//	var x int
//	var y float32
//	fmt.Println("请输入两个数：")
//	fmt.Scanln(&x,&y)
//	fmt.Println(x,y)
//	fmt.Println("请输入两个数：")
//	fmt.Scanf("%d,%f",&x,&y)
//
//}

//func main() {
//	var x interface{}
//
//	switch i := x.(type) {
//	case int:
//		fmt.Println("int",i)
//	case nil:
//		fmt.Println("nil",i)
//	case float32:
//		fmt.Println("float32",i)
//	}
//}

//func main() {
//	for i:=0;i<=10;i++{
//		if i > 5 {
//			break
//		}
//	}
//	for i:=0;i<=10;i++{
//		if i % 2 == 0 {
//			continue
//		}
//	}
//label:
//	fmt.Println("aaaaaaaa")
//	for i:=0;i<=10;i++{
//		if i > 5 {
//			goto label
//		}
//	}
//
//
//}
//func main() {
//
//	numbers := []int{1,2,3,4,5,5,6,7,8,9}
//	//
//	//fmt.Println(numbers)
//	//fmt.Println(numbers[1:])
//	//fmt.Println(numbers[:9])
//	//fmt.Println(numbers[2:3])
//	//fmt.Println(numbers[:])
//	var numbers1 = []int{}
//	numbers1 = numbers
//	fmt.Println(numbers1)
//	fmt.Println(append(numbers1,1))
//	fmt.Println(append(numbers1,numbers[1:2]...))
//	fmt.Println(copy(numbers,numbers1))
//	fmt.Println(numbers1)
//	fmt.Println(numbers)
//
//}

//func main() {
//	var mapNum  map[int]string
//	mapNum = make(map[int]string)
//	mapNum[1] = "test01"
//	mapNum[2] = "test02"
//
//	a,ok := mapNum[1]
//	fmt.Println(ok,a)
//
//	fmt.Println(len(mapNum))
//	delete(mapNum,1)
//	fmt.Println(mapNum)
//}

//func main() {

//str := "hello world"
//fmt.Println(str)
//for i := 0; i<len(str);i++ {
//	fmt.Println(str[i])
//}
//for i := 0;i<len(str);i++{
//	fmt.Printf("%c",str[i])
//}

//fmt.Println(strings.HasSuffix(str,"d"))
//fmt.Println(strings.HasPrefix(str,"h"))
//fmt.Println(strings.SplitN(str,"",1))
//fmt.Println(strings.Split(str,""))
//fmt.Println(strings.Count(str,"o"))
//fmt.Println(strings.Count(str,"o"))
//fmt.Println(strings.TrimSpace(str))
//fmt.Println(strings.Join(sj,"+"))
//fmt.Println(strings.Replace(str,"o","w",-1))

//var str = "123"
//fmt.Println(strconv.Atoi(str))
//fmt.Println(strconv.ParseInt(str,8,2))
//fmt.Println(strconv.Itoa(123))
//	test("12","23","34","45")
//
//}
//
//func test(arg ...string)  {
//	for i,v := range arg {
//		fmt.Println(i,v)
//	}
//}

//func main()  {
//	var a = map[string]int{
//		"test": 1,
//		"test2": 2,
//	}
//	test(a)
//	fmt.Println(a)
//}
//
//func test(n map[string]int)  {
//	fmt.Println(n)
//	n["test"] = 2
//	fmt.Println(n)
//}
//
//type test01 struct {
//	name string
//	id   int
//}
//
//type test02 struct {
//	test01
//	age int
//}
//
//func main() {
//
//	var t = test01{
//		name: "xielong",
//		id:   2,
//	}
//
//	var t2 = test02{
//		test01: t,
//		age:    18,
//	}
//
//	fmt.Println(t2.age)
//	fmt.Println(t2.id)
//	fmt.Println(t2.name)
//	fmt.Println(t2.test01.name)
//	fmt.Println(t2.test01.id)
//
//}

//type number1 struct {
//	name string
//	id int64
//}
//
//type number2 struct {
//	name map[string]int64
//}
//
//func main() {
//	n1 := number2{
//	name: map[string]int64{
//		"name": 1,
//	},
//	}
//	n2 := number2{
//		name: map[string]int64{
//			"name": 1,
//		},
//	}
//
//	if n1 == n2 {
//		fmt.Println("struct is equal")
//	} else {
//		fmt.Println("struct is not equal")
//	}
//}

//
//type number1 struct {
//	name string
//}
//
//type number2 struct {
//	number1
//}
//
//type number3 struct {
//	number1
//}
//
//func (n *number1)area()  {
//	fmt.Println("area")
//}
//
//func main() {
//	n1 := number1{name: "n1"}
//	n2 := number2{n1}
//	n3 := number2{n1}
//	n2.area()
//	n3.area()
//}
//
//type inter interface {
//	Call()
//}
//
//type number1 struct {
//	name string
//}
//
//type number2 struct {
//	name string
//}
//
//func (n *number2) Call() {
//	fmt.Println(n.name)
//}
//
//func (n *number1) Call() {
//	fmt.Println(n.name)
//}
//
//func main() {
//	var x inter
//	n1 := number1{
//		name: "number1",
//	}
//	x = &n1
//	x.Call()
//	n2 := number2{
//		name: "n2",
//	}
//
//	x = &n2
//	x.Call()
//
//	var i interface{} = number1{
//		name: "test",
//	}
//
//	s, ok := i.(number1)
//	if ok {
//		fmt.Println(s.name)
//	}
//
//	switch in := i.(type) {
//	case number1:
//		fmt.Println("number1111111",in.name)
//	}
//
//}

//func main() {
//	fileInfo,err := os.Stat("./test.txt")
//	if err != nil {
//		fmt.Println("open file err")
//	}
//	fmt.Println(fileInfo.Name())
//	fmt.Println(fileInfo.Size())
//	fmt.Println(fileInfo.IsDir())
//	fmt.Println(fileInfo.Mode())
//	fmt.Println(fileInfo.ModTime())
//	fmt.Println(fileInfo.Sys())
//}

//func main() {
//
//	fmt.Println(filepath.Abs("./test.txt"))
//	fmt.Println(filepath.Split("./test.txt"))
//	fmt.Println(filepath.Base("./test.txt"))
//	fmt.Println(filepath.Dir("./test.txt"))
//	fmt.Println(filepath.Join("D:/","test.txt"))
//	fmt.Println(filepath.IsAbs("./test.txt"))
//}

//func main() {
//
//	// 创建文件
//	file, err := os.Create("./test1.txt")
//	if err != nil {
//		fmt.Println("创建文件失败")
//	}
//
//	defer file.Close()
//	// 创建目录
//	err = os.Mkdir("test", 0666)
//	if err != nil {
//		fmt.Println("创建目录失败")
//	}
//
//	os.MkdirAll("/test/test1/test2", 0744)
//	//打开文件
//	filename, err := os.Open("./test.txt")
//	if err != nil {
//		fmt.Println("打开文件失败")
//	}
//
//	defer filename.Close()
//
//	fileINfo, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_APPEND, 0666)
//	if err != nil {
//		fmt.Println("open file failed")
//	}
//
//	defer fileINfo.Close()
//
//	os.Rename("./test.txt", "./testnew.txt")
//	os.Remove("./test1.txt")
//	os.RemoveAll("./test")
//
//}

//func main() {

//var count = make([]byte, 128)
//n, err := file.Read(count)
//if err != nil {
//	fmt.Println("read data failed")
//	return
//}
//
//if err == io.EOF {
//	fmt.Println("file read failed")
//	return
//}
//
//fmt.Println(string(count[:n]))
//	file, err := os.Open("./main.go")
//	if err != nil {
//		fmt.Println("open file failed")
//	}
//
//	defer file.Close()
//
//    var count []byte
//	var tmp = make([]byte,128)
//
//	for {
//		n, err := file.Read(tmp)
//		if err != nil {
//			fmt.Println("read data failed")
//			return
//		}
//
//		if err == io.EOF {
//			fmt.Println("read data over")
//			break
//		}
//
//		count = append(count,tmp[:n]...)
//	}
//	fmt.Println(string(count))
//}

//func main() {
//	file, err := os.Open("./test.txt")
//	if err != nil {
//		fmt.Println("open file failed")
//	}
//
//	defer file.Close()
//
//	reader := bufio.NewReader(file)
//	for {
//
//		str, err := reader.ReadString('\n')
//
//		if err == io.EOF {
//			fmt.Println("read data over")
//			break
//		}
//		if err != nil {
//			fmt.Println("read data failed")
//			return
//		}
//		fmt.Println(str)
//	}
//}

//func main() {
//	//file, err := os.Open("./test.txt")
//	//if err == nil {
//	//	fmt.Println("open file err")
//	//	return
//	//}
//	//
//	//defer file.Close()
//	str, err := ioutil.ReadFile("./test.txt")
//	if err != nil {
//		fmt.Println("reader data failed")
//	}
//	fmt.Println(string(str))
//}
//
//func main() {
//	//file, err := os.OpenFile("./test1.txt",os.O_APPEND|os.O_CREATE,0666)
//	//if err != nil {
//	//	fmt.Println("open file err")
//	//	return
//	//}
//	//
//	str := "hello world"
//
//	//_, err = file.Write([]byte(str))
//	//if err != nil {
//	//	fmt.Println("write data err")
//	//	return
//	//}
//	//_, err = file.WriteString(str)
//	//if err != nil {
//	//	fmt.Println("write data err ")
//	//	return
//	//}
//
//	//reader := bufio.NewWriter(file)
//	//reader.WriteString(str)
//	//reader.Flush()
//	ioutil.WriteFile("./test1.txt", []byte(str), 0666)
//	ioutil.
//}

//func main() {
	//dir , err := ioutil.ReadDir("D:/go/src")
	//if err != nil {
	//	fmt.Println("open dir failed")
	//	return
	//}
	//for _, v := range dir {
	//	fmt.Println(v.Name())
	//	v.IsDir()
	//}
//	listDir("D:/go/src",0)
//}

//func listDir(file string, level int)  {
//	s := "|--"
//	for i:=0;i<level;i++{
//		s = "| " +s
//	}
//
//	fileInfo, err := ioutil.ReadDir(file)
//	if err != nil {
//		fmt.Println("open file failed")
//		return
//	}
//
//
//	for _,v := range fileInfo{
//		fileName := file + "/" + v.Name()
//		fmt.Printf("%s%s\n",s,fileName)
//		if v.IsDir() {
//			level = level + 1
//			listDir(fileName,level)
//		}
//	}
//}













