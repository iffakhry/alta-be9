package main

import "fmt"

func palindrome(input string) bool {
	var isPalindrome bool
	var strTemp string

	//mendapatkan panjang string harus menggunakan rune
	var strLength int = len([]rune(input))

	//string terbentuk dari byte, sehingga tidak bisa langsung di len untuk dapat panjang string, karena nanti yang dibaca akan byte nya.
	// fmt.Println([]byte(input))

	for i := strLength - 1; i >= 0; i-- {
		strTemp += string(input[i])
	}

	if strTemp == input {
		isPalindrome = true
	}

	return isPalindrome
}

func palindrome2(input string) bool {
	var st, en int
	en = len(input) - 1
	for en >= st {
		fmt.Println("val ", input[st], input[en-st])
		if input[st] != input[en-st] {
			return false
		}
		st++
	}
	return true
}

func main() {
	fmt.Println(palindrome(("kupu-kupu")))
	fmt.Println(palindrome(("kakak")))
	fmt.Println(palindrome2(("kupu-kupu")))
	fmt.Println(palindrome2(("noon")))
}
