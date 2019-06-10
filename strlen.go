package piscine

import "fmt"

func StrLen(str string) int{
	nb:=0
	for _,n := range str {
		nb++
	}
	return nb
}
