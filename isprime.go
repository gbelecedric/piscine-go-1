package piscine

func IsPrime(nb int) bool{
	decision:=true
	for i:=2;i<nb;i++{
		if nb%i==0{
			decision=false
		}
	}
	return decision
}



