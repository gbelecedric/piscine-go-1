package piscine

func FindNextPrime(nb int)int{
	nextprime=nb
	for i>nb{
		nextprime++
		if premier(nextprime){
			return nextprime
		}
		i++	
	}

}


func premier(nb int) bool{
	decision:=true
	if nb<=1{
		return false
	}
	for i:=2;i<nb;i++{
		if nb%i==0{
			decision=false
		}
	}
	return decision
}



