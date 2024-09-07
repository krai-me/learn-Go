package main //package declare package main runs func main(){}
import "fmt" //allows us to use code from other packages
// We use the func keyword to declare the Go function main:
// func main(){} is main entry point of go file
func main() {
	fmt.Println("hello world") //to print
	fmt.Println(`this is
		mutlti line
		print`) //for multiline we use backticks
}
