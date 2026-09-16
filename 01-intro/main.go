package main

import (
	"fmt"
	"os"

	"github.com/AlexeyD1982/mymath"
	mymathV2 "github.com/AlexeyD1982/mymath/v2"
	"github.com/mdp/qrterminal/v3"
)

func main() {
	config := qrterminal.Config{
		Level:     qrterminal.M,
		Writer:    os.Stdout,
		BlackChar: qrterminal.WHITE,
		WhiteChar: qrterminal.BLACK,
		QuietZone: 1,
	}

	qrterminal.GenerateWithConfig("Hello, world", config)
	fmt.Println(mymath.Add(1, 2))
	fmt.Println(mymath.Sub(4, 2))
	fmt.Println(mymathV2.Add(1, 2, 3))
}
