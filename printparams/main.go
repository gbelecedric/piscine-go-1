package main

import (
	"fmt"
	"os"
	//"strings"
	)

func main(){
	args:=os.Args
	//cp:=strings.Split(args[1]," ")
	if len(args)> 1{
		for i:=1;i<len(args);i++{
			fmt.Printf("%v\n",args[i])
		}
		
	}
}
