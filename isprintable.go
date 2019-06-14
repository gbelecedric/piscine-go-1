package piscine

func IsPrintable(str string) bool{
	cbstr:=[]rune(str)
	for i,_:=range cbstr{	
		if !(cbstr[i]>=65 && cbstr[i]<=90){
			return false
		}
	}
	
	return true
}
