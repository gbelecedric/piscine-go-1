package piscine

import "strings"

func Capitalize(s string) string{
	cbs:=[]rune(s)
	for i,_:=range cbs{
		if(i==0){
			if(cbs[i]>=97 &&cbs[i] <=122){//on met la premier lettre en majuscule
				cbs[i]-=32
			}
			
		}else{
			if(cbs[i]>=65 &&cbs[i] <=90){//on met les autres lettres en miniscule
				cbs[i]+=32
			}
		}
		
	}

	return string(cbs)
}
