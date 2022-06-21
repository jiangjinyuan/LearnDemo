package main

import (
	"fmt"
	"unsafe"
)

type Coder interface {
	code()
}

type Gopher struct {
	name string
}

func (g Gopher) code() {
	fmt.Printf("%s is coding\n", g.name)
}

type MyError struct{}

// 实现 error 接口
func (i MyError) Error() string {
	return "MyError"
}

// Process 函数返回了一个 error 接口，这块隐含了类型转换。
func Process() error {
	var err *MyError = nil
	return err
}

func CompareInterfaceValueAndNil() {
	var c Coder
	// true
	fmt.Println(c == nil)
	// c: <nil>, <nil>
	fmt.Printf("c: %T, %v\n", c, c)

	var g *Gopher
	// true
	fmt.Println(g == nil)

	// g赋值给c后，c的动态类型为*Gopher，值为<nil>
	c = g
	// false
	fmt.Println(c == nil)
	// c: *main.Gopher, <nil>
	fmt.Printf("c: %T, %v\n", c, c)

	err := Process()
	// <nil>
	fmt.Println(err)
	// false
	fmt.Println(err == nil)
	// 如何检测 MyError 类型是否实现了 error 接口？
	// 下面两种方法均可，如果 MyError 没有实现了 Error() 方法，编译的时候即会出错；
	var _ error = MyError{}
	var _ error = (*MyError)(nil)
}

type iface struct {
	itab, data uintptr
}

func PrintInterface() {
	var a interface{} = nil

	var b interface{} = (*int)(nil)

	x := 5
	var c interface{} = (*int)(&x)

	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))
	ic := *(*iface)(unsafe.Pointer(&c))

	fmt.Println(ia, ib, ic)

	fmt.Println(*(*int)(unsafe.Pointer(ic.data)))
}

type Student struct {
	Name string
	Age  int
}

// 增加对String()方法的实现，fmt.Println函数就可以按照我们自定义的方法来打印;
func (s Student) String() string {
	return fmt.Sprintf("[Name: %s], [Age: %d]", s.Name, s.Age)
}

func TypeConversionAndAssertion() {
	var i interface{} = new(Student)
	// 安全的类型断言
	s, ok := i.(Student)
	if ok {
		fmt.Println(s)
	}
	// switch 类型断言
	switch v := i.(type) {
	case nil:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("nil type[%T] %v\n", v, v)

	case Student:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("Student type[%T] %v\n", v, v)

	case *Student:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("*Student type[%T] %v\n", v, v)

	default:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("unknow\n")
	}
	// 引申1
	// fmt.Println 函数的参数是 interface。对于内置类型，函数内部会用穷举法，得出它的真实类型，然后转换为字符串打印。
	// 而对于自定义类型，首先确定该类型是否实现了 String() 方法，如果实现了，则直接打印输出 String() 方法的结果；
	// 否则，会通过反射来遍历对象的成员进行打印。
	var a = Student{
		Name: "qcrao",
		Age:  18,
	}

	// fmt.Println(a)

	// 增加对String()方法的实现，fmt.Println函数就可以按照我们自定义的方法来打印;
	fmt.Println(a)
}

func main() {
	// 接口类型和 nil 作比较
	CompareInterfaceValueAndNil()
	// 打印出接口的动态类型和值
	PrintInterface()
	// 类型转换和类型断言
	TypeConversionAndAssertion()
}
