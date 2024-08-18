package main

import (
	"fmt"
	"slices"
	"unsafe"
)

func main() {

	// 3 examples to show that when a slice grows beyond its capacity the data moves to another location.
	// Companion to https://badfutures.github.io/2024/08/18/3.GoLangSlices.html
	example1()
	example2()
	example3()
}

func example1() {
	fmt.Println()
	fmt.Printf("Example 1. Creating slice by specifying capacity explicitly \n")

	s1 := make([]int, 0)
	print("s1", &s1)
	for i := 0; i < 5; i++ {
		s1 = append(s1, i)
		print("s1", &s1)
	}

	s2 := s1

	fmt.Println()
	fmt.Printf("s2: = s1\n")

	print("s1", &s1)
	print("s2", &s2)

	fmt.Println()
	fmt.Printf("s1 = slices.Delete(s1, 1, 3)\n")
	s1 = slices.Delete(s1, 1, 3)

	print("s1", &s1)
	print("s2", &s2)

	fmt.Println()
	fmt.Printf("s1[0:3:3]\n")
	s3 := s1[0:3:3]

	print("s1", &s1)
	print("s2", &s2)
	print("s3", &s3)

	fmt.Println()
	fmt.Printf("s1[0]=10, s2[1] = 11; s3[2] = 12\n")
	s1[0] = 10
	s2[1] = 11
	s3[2] = 12

	print("s1", &s1)
	print("s2", &s2)
	print("s3", &s3)

	// Let's grow s3

	s3 = append(s3, 15)
	fmt.Println()
	fmt.Printf("s3 = append(s3, 15)\n")

	print("s1", &s1)
	print("s2", &s2)
	print("s3", &s3)

	fmt.Println()
	fmt.Printf("s1[0]=20, s2[1] = 21; s3[2] = 22\n")
	s1[0] = 20
	s2[1] = 21
	s3[2] = 22

	print("s1", &s1)
	print("s2", &s2)
	print("s3", &s3)

}

func example2() {

	fmt.Println()
	fmt.Printf("Example 2. Appending after len = cap means getting a new buffer\n")

	s1 := make([]int, 5)
	for i := 0; i < 5; i++ {
		s1[i] = i
	}
	s2 := s1
	print("s1", &s1)
	print("s2", &s2)

	fmt.Println()
	fmt.Printf("Append 1 elements to s1 when len = cap\n")

	s1 = append(s1, 6)
	print("s1", &s1)
	print("s2", &s2)

	fmt.Println()
	fmt.Printf("s1[0] = 10\n")
	s1[0] = 10
	print("s1", &s1)
	print("s2", &s2)

}

func example3() {
	fmt.Println()
	fmt.Printf("Example 3, growing within capacity doesn't copy.\n")

	s1 := make([]int, 0)
	for i := 0; i < 5; i++ {
		s1 = append(s1, i)
		print("s1", &s1)
	}

	fmt.Println()
	fmt.Printf("s2 = s1\n")
	fmt.Printf("slices.Delete(s1, 2, 5)")
	s2 := s1
	s1 = slices.Delete(s1, 2, 5)

	print("s1", &s1)
	print("s2", &s2)

	fmt.Println()
	fmt.Printf("Append 1 elements to s1 when len = cap\n")

	for len(s1) < cap(s2) {
		s1 = append(s1, 3)
		print("s1", &s1)
		print("s2", &s2)
	}

	fmt.Println()
	fmt.Printf("s1[0] = 10\n")
	s1[0] = 10
	print("s1", &s1)
	print("s2", &s2)

	fmt.Printf("append(s1, 4)\n")
	s1 = append(s1, 4)
	print("s1", &s1)
	print("s2", &s2)

	fmt.Println()
	fmt.Printf("s1[0] = 11\n")
	s1[0] = 11
	print("s1", &s1)
	print("s2", &s2)
}

func print(name string, s *[]int) {
	if len(*s) > 0 {
		fmt.Printf("%s=%v len=%d, cap=%d, %v\n", name, *s, len(*s), cap(*s), (unsafe.Pointer(&(*s)[0])))
	} else {
		fmt.Printf("%s=%v len=%d, cap=%d\n", name, *s, len(*s), cap(*s))
	}
}
