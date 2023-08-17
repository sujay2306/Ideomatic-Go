
// Variadic functions can be called with any number of trailing arguments.
//  For example, fmt.Println is a common variadic function.

package main
import "fmt"

func sum (nums ...int) int {  //... indicates variadics
	sum := 0
	for _,n := range nums {
		sum += n
	}
	return sum
}

func main(){
	a := [] int {1,1,2,3,4,5}
	b := [] int {15,6,7,8,8,8}
	all := append(a,b...) //contains both a and b
	fmt.Println(all)
	
	answer := sum(all...)
	fmt.Println(answer)

}
