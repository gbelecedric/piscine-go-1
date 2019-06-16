package piscine

func MakeRange(min, max int) []int{
	result:=make([]int,max-min)
	if max<=min{
		return nil
	}
	j:=0
	for i:=min;i< max;i++{
		result[j]=i
		j++
	}
	return result
}
