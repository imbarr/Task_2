package main

import (
	"Task_2/matrix"
	"os"
)

func main() {
	_, e := matrix.ReadMatrix(os.Stdin)
	print(e.Error())
}
