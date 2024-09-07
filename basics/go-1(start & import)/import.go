//to import multiple packages we can do
/*
import "package1"
import "package2"
*/
//or
/*
import(
	"package1"
	"package2
)
*/
//for alias
/*
import(
	p1 "package1"
	p2 "package2"
	"package3"
)
*/

//so we call func by p1.SampleFunc()
//instead of package1.SampleFunc()

package main
import(
	"fmt"
	t "time"
)
func main(){
	fmt.Println("Hello World")
	fmt.Println(t.Now())// to get time
}
