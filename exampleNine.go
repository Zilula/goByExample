package main 
import "fmt"

func main() {

	s := make([]string, 3)

	fmt.Println("empty array with three strings", s)

	s[0] = "one"
	s[1] = "two"
	s[2] = "three"

	fmt.Println("set", s)
	fmt.Println("get", s[2])
	fmt.Println("length", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("get index 5", s[5])
	fmt.Println("appended some stuff", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copyed S to make C", c)
	l := s[2:5]
	fmt.Println("created L with the indexs of 2-5 from s", l)

	l = s[:5]
	fmt.Println("gets evertyhing up to and including index 5 from s array", l)

	l = s[2:]
	fmt.Println("gets everything after index 2 from s array", l)

	t := []string{"g", "h", "i"}
	fmt.Println("declated array with strings", t)

	twoD := make([][]int, 3)
	fmt.Println("Step 1 make a 2D array", twoD)
	for i := 0; i < 3; i ++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		fmt.Println("step two make each index another array", twoD)
		for j := 0; j < innerLen; j ++ {
			twoD[i][j] = j + i
		}
	}
	fmt.Println("2d: ", twoD)



}