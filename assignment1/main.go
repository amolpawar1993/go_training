package main

import "fmt"

func main(){
	fmt.Println("Star Pattern")
	star()
	fmt.Println("Fibonacci Series")
	fibbonaci()
	fmt.Println("Reverse print")
	val:=reverse("123")
	print(val)
	fmt.Println("\nNumbers Pattern")
	num_pattern()

}

func star() {
	for i:=0; i<=4;i++{
		for j:=0;j<=i;j++{
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func fibbonaci()  {
	f1 := 0
	f2 := 1
	fmt.Print(f1," ",f2," ")
	for i:=2; i<=10;i++{
		next:=f1+f2
		fmt.Print(next, " ")
		f1 = f2
		f2 = next
	}
}

func reverse(s string) (result string){
	for _, v:= range s{
		result = string(v) + result
	}
	return
}

func num_pattern(){
	n := 1
	for i:=1; i<5;i++{
		for j:=0;j<i;j++{
			fmt.Print(n," ")
			n++
		}
		fmt.Println()
	}
}