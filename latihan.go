// Deret Geometri
package main
import "fmt"

func geo(a, r, n int)int {
	if n == 1{
		return a
	}
	return a * geo(r, r, n-1)
}
func main(){
	var a, r, n int
	fmt.Scan(&a, &r, &n)
	fmt.Print(geo(a, r, n))
}