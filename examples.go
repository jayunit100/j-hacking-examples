package main
import (
	"fmt"
	"strings"
	"strconv"
)
// How do you read a struct from a line of text?
type Example1 struct {
	aString string
	aNum int64
}

// input: a string
// output: a struct.
func example1_stringToStruct(in string) Example1 {
	slice := StrSlice(strings.Split(in, ","))
	aString, _:= slice.Get(0)
	aNum, _:= strconv.ParseInt(slice.Get(1),10,64)
	// We can create a struct by relying on field ordering.
	return Example1{
		aString,
		aNum,
	}
}

// input: A map of ints->string
// output: A string array, where the map key == the index
func example2_mapToArray(m1 map[string]int) {
	m := map[int]string{
		1 : "a", 2 : "b",
	}
	v := make([]string, 0, len(m))

	for _, value := range m {
		v = append(v, value)
	}
}

// output 1.0000 e+09
func example3_sciNot(i float64) string {
	return fmt.Sprintf("sci: %e",i)
}

func main() {
	// fmt.Println(fmt.Sprintf(example1_stringToStruct("hello,100")))
	fmt.Println(example1_stringToStruct("hello,100"))
	fmt.Println(example3_sciNot(1000000000))
}

// less intersting utilities...
// Here's how to add a function to a primitive.
type StrSlice []string
func (s StrSlice) Get(i int) string {
	if i >= 0 && i < len(s) {
		return s[i]
	}
	return ""
}