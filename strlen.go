package piscine

func StrLen(str string) int{
	nb:=0
	for _,n := range str {
		fmt.Printf("%c",n)
		nb++
	}
	return nb
}
