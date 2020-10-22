package main

import (
        "bufio"
        "fmt"
        "log"
        "os"
        "strconv"
        "strings"
	"mathcalculator"
	)

func input() int {
        reader := bufio.NewReader(os.Stdin)
        in, err := reader.ReadString('\n')
        if err != nil {
                log.Fatal(err)
       }
in = strings.TrimSpace(in)
number, err := strconv.Atoi(in)
if err != nil {
        log.Fatal(err)}
	return number
}

func absolute(n int) int {
        if n < 0 {
	        return n * -1
	}
        return n
}

func factorial(n int) int {
        if n == 0 {
                return 1
	}
        return n * factorial(n-1)
}
															func fibonacci(n int) int {
        if n == 1 || n == 2 {
                return 1
        } else {
	        return fibonacci(n-1) + fibonacci(n-2)
        }
}

func power(b int, e int) int{
        r := 1
        for i := 1; i<=e ; i++{
                r = r * b
        }
        return r
}

func division (a int, b int) int{
        if b == 0{
	        fmt.Print("0을 제외한 숫자를 입력하세요")
	}
	return a/b

func oddeven (n int) int{
	if n % 2 == 0{
		fmt.Println(n,"은 짝수입니다.")
	}
	else{
		fmt.Println(n, "은 홀수입니다.")
	}
func permutation(n int, r int) int {
	if n == 0{
		return 1
	}
	return  n * circular(n-1) / (n-r) *circular(n-r-1)

func main() {
	for x := 1; true; x++ {
		fmt.Print("1) 절대값 2) 팩토리얼 3) 피보나치 4) 거듭제곱 5)나눗셈 6) 7)순열 8)소수판정 9)조합 10) 종료 : ")
		n := input()
		if n == 1 {
			fmt.Print("정수 입력(절대값 계산) : ")
			fmt.Println(absolute(input()))
		} else if n == 2 {
			fmt.Print("정수 입력(팩토리얼 계산) : ")
			fmt.Println(factorial(input()))
		} else if n == 3 {
			fmt.Print("정수 입력(피보나치 출력) : ")
			f := input()
			fmt.Println(fibonacci(f))
		} else if n == 4 {
			fmt.Print("Base 입력 : ")
			b := input()
			fmt.Print("Exponent 입력 : ")
			e := input()
			fmt.Println(power(b, e))
		} else if n == 5 {
			fmt.Print("정수 입력(나누기 계산) :")
			a := input()
			b := imput()
			fmt.Println(division(a,b))
		} else if n == 6 {
			fmt.Print("정수 입력(홀수 짝수 구분):")
			fmt.Print(oddeven(input()))
		} else if n == 7 {
			fmt.Print("정수 입력(순열계산)")
			r := input()
			fmt.Print(permutation(input(),r))
		} else if n == 8 {
			fmt.Print("정수 입력 (소수 계산)")
			mathcalculator.Prime(input())
		} else if n == 9 {
			fmt.Print("정수 입력(조합 계산)")
			r := input()
			mathcalculator.Combination(input(),r)
		} else if n == 10{
			os.Exit(3)
		}
		else {
			fmt.Print("잘못 된 입력 값입니다. 1~10사이의 수를 입력하시오.")

		}

	}

}
