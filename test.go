package piscine

/*wdmatch
func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	valid := false
	valid1 := false
	valid2 := false
	for i := 0; i < 3; i++ {
		if args[0][0] == args[1][i] {
			valid = true
		}
	}
	for i := len(args[1]) - 1; i >= len(args[1])-3; i-- {
		if args[0][len(args[0])-1] == args[1][i] {
			valid1 = true
		}
	}
	for i := 1; i < len(args[0])-1; i++ {
		for j := len(args[1])-1; j > len(args[1])-3; j-- {
			if args[0][i] == args[1][j] {
				valid2 = true
			}
		}
		if !valid2 {
			fmt.Println("error1")
			return
		}
	}
	if valid && valid1 {
		fmt.Println(args[0])
	}
}
*/

/* ispowerof2
func main() {
	args := os.Args
	res, _ := strconv.Atoi(args[1])
	if t(res) {
		fmt.Println("True")
	} else {
		fmt.Println("False")

	}
}

func t(n int) bool {

	for n != 2 {
		if n%2 != 0 {
			return false
		}
		n /= 2
	}
	return true
}
*/

/* searchreplace
func main() {
	args := os.Args
	res := ""
	if len(args) != 4 {
		return
	}
	for i := 0; i<len(args[1]); i++ {
		if string(args[1][i]) == args[2] {
			res += string(args[3])
		} else {
			res += string(args[1][i])
		}
	}
	fmt.Println(res)
}
*/

/* doop
func main() {
	args := os.Args
	if len(args) != 4 {
		return
	}
	if doop(args) == "" {
	fmt.Print(doop(args))
	} else {
	fmt.Println(doop(args))
	}
}
func aToi(str string) int {//125
	res := 0

	t := 1
	for i := 0; i < len(str); i++ {
		if str[i] == '-' && i == 0 {
			t = -1
		}
		if str[i] >= '0' && str[i] <= '9' {
			res = res*10 + int(str[i]-'0')
		}
	}
	return res * t
}

func iToa(nb int) string {//125
	res := ""
	for nb != 0 {
		res = string(nb%10+'0') + res
		nb /= 10
	}
	return res
}

func doop(args []string) string {
	nb1 := aToi(args[1])
	nb2 := aToi(args[3])
	signe := args[2]
	if !((nb1 >= 0 && nb1 <= 9) || (nb2 >= 0 && nb2 <= 9)) {
		return ""
	}
	if signe == "+" {
		return iToa(nb1 + nb2)
	} else if signe == "-" {
		return iToa(nb1 - nb2)
	} else if signe == "/" {
		if nb2 == 0 {
			return "No division by 0"
		}
		return iToa(nb1 / nb2)
	} else if signe == "%" {
		if nb2 == 0 {
			return "No modulo by 0"
		}
		return iToa(nb1 % nb2)
	} else if signe == "*" {
		return iToa(nb1 * nb2)
	} else {
		return ""
	}
}
*/

/* IsCapitalized
func IsCapitalized(s string) bool {
	g := []rune(s)
	if len(s) == 0 || g[0] >= 'a' && g[0] <= 'z' {
		return false
	}
	for i := 0; i < len(s); i++ {
		if g[i] == ' ' {
			if g[i+1] >= 'a' && g[i+1] <= 'z' {
				return false
			}
		}
	}
	return true
}
*/

/* PrintMemory
func PrintMemory(arr [10]byte) {
	base := "0123456789abcdef"

	for i := 0; i < len(arr); i++ {
		div := int(arr[i]) / len(base)
		mod := int(arr[i]) % len(base)
		z01.PrintRune(rune(base[div]))
		z01.PrintRune(rune(base[mod]))
		if i != 3 && i != 7 && i != 9 {
			z01.PrintRune(' ')
		}
		if i == 3 || i == 7 || i == 9 {
			z01.PrintRune('\n')
		}
	}
	for i := 0; i < len(arr); i++ {
		if unicode.IsGraphic(rune(arr[i])) {
			z01.PrintRune(rune(arr[i]))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}
*/

/* union
func main() {
	input := os.Args[1] + os.Args[2]
	arr := make(map[string]int)
	res := make(map[string]bool)

	for _, char := range input {
		arr[string(char)]++
	}
	fmt.Println(arr)

	for _, char := range input {
		if !res[string(char)] {
			fmt.Printf(string(char))
			res[string(char)] = true
		}
	}
	fmt.Println()
}
*/

/* inter

func main() {
	if len(os.Args) != 3 {
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]

	seen := make(map[rune]bool)
	result := ""

	for _, char := range str1 {
		if !seen[char] {
			seen[char] = true
			result += string(char)
		}
	}

	for _, char := range str2 {
		if seen[char] {
			fmt.Print(string(char))
			seen[char] = false
		}
	}
	fmt.Println()
}
*/

/* expandstr
func main() {
	args := os.Args[1:]
	count := 0
	res := ""
	isFirst := true
	for i := 0; i < len(args[0]); i++ {
		if args[0][i] != ' ' {
			isFirst = false
		}
		if args[0][i] == ' ' {
			count++
			continue
		} else if count != 0 && isFirst == false {
			res += "   " + string(args[0][i])
			count = 0
		} else if count == 0 {
			res += string(args[0][i])
		}
	}
	fmt.Println(res)
}
*/

/* reversestrcap
func main() {
	args := os.Args[1:]
	str := ""
	for _, c := range args {
		for _, v := range c {
			if v >= 'A' && v <= 'Z' {
				v += 32
			}
			str += string(v)
		}
		str += "\n"
	}
	myString := ""
	for i := 0; i < len(str); i++ {
		if i == len(str)-1 {
			break
		}
		if str[i+1] == ' ' {
			if str[i] >= 'a' && str[i] <= 'z' {
				myString += string(str[i] - 32)
			} else {
				myString += string(str[i])

			}
		} else {
			if str[i+1] == '\n' {
				if str[i] >= 'a' && str[i] <= 'z' {
					myString += string(str[i] - 32)
				} else {

					myString += string(str[i])
				}
			} else {
				myString += string(str[i])

			}
		}

	}
	fmt.Println(myString)
}


func AtoiBase(str string, base string) int {

}
*/
