package main

import "fmt"

func generateParenthesis(n int) []string {
	return generateParenthesisR(make([]string, 0), n, "")
}

func generateParenthesisR(strs []string, n int, s string) []string {
	if len(s) > n*2 {
		return strs
	} else if len(s) == n*2 {
		if isValid(s) {
			return []string{s}
		}
		return []string{}
	}
	stack1 := generateParenthesisR(strs, n, s+"(")
	stack2 := generateParenthesisR(strs, n, s+")")
	strs = append(strs, stack1...)
	strs = append(strs, stack2...)

	return strs
}

func isValid(s string) bool {
	stack := make([]string, len(s))
	top := 0
	for i := 0; i < len(s); i++ {
		if top > -1 && string(s[i]) == "(" {
			stack[top] = string(s[i])
			top++
		} else if string(s[i]) == ")" {
			top--
		}
	}
	return top == 0
}

func main() {
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(4))
}

/*
//							idea:
func generateFromPrefix(o int, c int, prefix string, results *[]string) {
    if o > 0 {
        newPrefix := prefix + "("
        generateFromPrefix(o-1, c, newPrefix, results)
    }
    if o < c {
        newPrefix := prefix + ")"
        generateFromPrefix(o, c-1, newPrefix, results)
    }
    if o == 0 && c == 0 {
        *results = append(*results, prefix)
    }
}

func generateParenthesis(n int) []string {
    answer := []string{}
    generateFromPrefix(n, n, "", &answer)
    return answer
}

//						very hard to understand
*/
