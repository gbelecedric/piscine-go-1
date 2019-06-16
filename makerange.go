package main
import "fmt"
func main(){
	fmt.Println(MakeRange(14,10))
}

func MakeRange(min, max int) []int{
	if max<=min{
		return nil
	}
	result:=make([]int,max-min)
	j:=0
	for i:=min;i< max;i++{
		result[j]=i
		j++
	}
	return result
}
