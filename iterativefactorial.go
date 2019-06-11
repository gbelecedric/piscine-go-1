package piscine

func IterativeFactorial(n int) int{
	if n == 0 || n==1{
		return 1
	}
	if n > 1{
		fact:=1
		for i:=1;i<=n;i++{
			fact*=i	
		}
		return fact
	}
	return 0
}


