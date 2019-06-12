package piscine

import "math"

func IterativePower(nb,power int) int{
	if n < 0{
		return 0
	}
	if n >=0{
		result:=1
		for i:=1;i<=nb;i++{
			result*=nb
			if result > math.MaxInt32{
				return 0			
			}
		}
		return result
	}
}


