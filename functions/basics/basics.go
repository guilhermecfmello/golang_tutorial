package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s\n", p1, p2)
}

func f5() (string, string) {
	return "first return", "second return"
}

func f6() (string, string, string) {
	return "1", "2", "3"
}

func main() {
	f1()
	f2("Gui", "Mello")
	fmt.Println(f4(f3(), "test"))
	s1, s2 := f5()
	fmt.Printf("F5: %s, %s\n", s1, s2)

	f1, _, _ := f6()
	fmt.Printf("F6: %s\n", f1)
}
