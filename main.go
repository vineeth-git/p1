package uniq
import "fmt"
import 	"github.com/StudioSol/set"

func main() {
	duplicatedInt64 := []int64{1, 1, 2, 2, 3, 3}
	fmt.Println(UniqueInt(duplicatedInt64))
}
func UniqueInt(input []int64) []int64{
	unduplicatedInt64 := set.NewLinkedHashSetINT64(input...)
	unduplicatedArray := unduplicatedInt64.AsSlice()
	return unduplicatedArray
}