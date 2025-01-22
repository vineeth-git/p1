package main
import "fmt"
import "github.com/vineeth-git/p1/uniq"

func main() {
	duplicatedInt64 := []int64{1, 1, 2, 2, 3, 3}
	fmt.Println(uniq.UniqueInt(duplicatedInt64))
}
