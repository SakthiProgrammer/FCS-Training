package basics

func Palindrome(str string) bool{
	left , right := 0, len(str) - 1

	for left < right {

		if(str[left] != str[right]){
			return false
		}

		left++
		right--
	}

	return true

}