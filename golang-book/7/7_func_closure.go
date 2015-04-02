package main

/* Task: Using makeEvenGenerator as an example, write a makeOddGenerator function that generates odd numbers. */
import "fmt"

func makeOddGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		if i%2 == 0 {
			i += 1
			ret = i
		} else {
			i += 2
			ret = i
		}
		return
	}
}
func main() {
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd()) // 1
	fmt.Println(nextOdd()) // 3
	fmt.Println(nextOdd()) // 5
}
