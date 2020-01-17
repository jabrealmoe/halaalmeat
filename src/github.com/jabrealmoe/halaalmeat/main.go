package main
import "fmt"

func say_hello(msg string){
	fmt.Println(msg);
}

func return_anon_func() func(string){
	// return an anonymous function
	return func(msg string){
		fmt.Print(msg)
	}
}

func main(){
	say_hello("say Hello")
	// anonyhmous function declared
	func(msg string){
            fmt.Println(msg)
        }("Hello Anonymous")

	print_func := return_anon_func()
	other_func := return_anon_func()
	other_func("funk u up!")
	print_func("Hello returned from anon")

}
