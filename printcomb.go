package piscine

import "fmt"

func PrintComb(){
	for i:=0; i <= 9; i++{
		for j:=i+1; j <= 9; j++{
			for k:=j+1; k <= 9;k++{
				fmt.Println(i,j,k)
			}
		}
	}
}
