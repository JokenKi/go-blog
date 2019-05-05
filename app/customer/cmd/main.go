package main

import (
	"flag"
	"fmt"
)

var flagvar int
var confPath string

func init() {
	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
	flag.StringVar(&confPath, "conf", "", "default config path")
}

func main(){
	flag.Parse()
	//flag.StringVar(&confPath, "conf", "", "default config path")
	var ip = flag.Int("flagname2", 1234, "help message for flagname")
	fmt.Println("ip has value ", *ip)
	fmt.Println("flagvar has value ", flagvar)
	fmt.Println("confPath has value ", confPath)
	
}