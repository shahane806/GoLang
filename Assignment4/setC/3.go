// wap in go language to demonstrate working embedded interface

package main

import "fmt"

type Cat interface {
	mew()
}
type Dog interface {
	bark()
}
type Animal interface {
	Cat
	Dog
}
type cat struct{}
type dog struct{}
type animal struct{}
func (c *cat) mew() {
	fmt.Println("Mew...")
}
func (d *dog) bark() {
	fmt.Println("BowBow...")
}
func (a *animal) hello(){
	fmt.Println("I am animal")
}
func main() {
	var i Cat
	var j Dog
	var cat1 cat
	var dog1 dog
	
	var n animal
	i = &cat1
	j = &dog1
	

	n.hello()
	i.mew()
	j.bark()


	


}
