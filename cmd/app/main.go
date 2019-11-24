package main

import (
	"github.com/wolfgarnet/salary"
)

func main() {
	system := salary.Initialize()
	salary.RunServer(system)
}
