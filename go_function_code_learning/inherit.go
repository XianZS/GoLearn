package main

import "fmt"

/*
Go 语言没有传统面向对象语言中的类(class)和继承(inheritance)概念，
而是通过组合(composition)和接口(interface)来实现类似的功能。
*/

type IMake interface {
	Make()
}

type Animal struct {
	name string
}

func (animal *Animal) Say() {
	fmt.Printf("I am an %v\n", animal.name)
}

type Dog struct {
	Animal
	age int
}

func (dog *Dog) Make() {
	dog.Animal.Say()
	fmt.Println("wang wang wang!")
}

type Cat struct {
	Animal
	age int
}

func (cat *Cat) Make() {
	cat.Animal.Say()
	fmt.Println("miao miao miao!")
}

func main() {
	/*
		=== [1]类"继承"的简单实现：组合(composition) ===
		在这里先创建Animal结构体，然后在创建Dog和Cat结构体时，
		将Animal结构体作为其字段，这样就可以实现继承的效果。
		其中，Animal作为父类，Dog和Cat作为子类。
		当子类想要调用父类的方法时，只需要通过子类的字段来调用即可。
		如:cat.Animal.Say()来达到效果。
	*/
	dog := Dog{
		Animal: Animal{name: "animal"},
		age:    10,
	}
	dog.Make()
	cat := Cat{
		Animal: Animal{name: "cat"},
		age:    20,
	}
	cat.Make()
	fmt.Println("---------")
	/*
		=== [2]类"继承"的简单实现：接口(interface) ===
		在这里先创建IMake接口，然后在创建Dog和Cat结构体时，
		将IMake接口作为其字段，这样就可以实现继承的效果。
		其中，IMake作为父类，Dog和Cat作为子类。
		当子类想要调用父类的方法时，只需要通过子类的字段来调用即可。
		如:cat.IMake.Make()来达到效果。
		注意：在Go语言中，接口是隐式实现的，也就是说，
		子类不需要显式地实现父类的接口，只要子类的方法签名和父类的接口方法签名一致，
		就可以认为子类实现了父类的接口。
	*/
	var iMakeDog, iMakeCat IMake
	iMakeDog, iMakeCat = &dog, &cat
	iMakeDog.Make()
	iMakeCat.Make()
}
