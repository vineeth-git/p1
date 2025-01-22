package uniq
import 	"github.com/StudioSol/set"
func UniqueInt(input []int64) []int64{
	unduplicatedInt64 := set.NewLinkedHashSetINT64(input...)
	unduplicatedArray := unduplicatedInt64.AsSlice()
	return unduplicatedArray
}