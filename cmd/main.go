package main

import (
	"github.com/billzayy/cli-financial-go/pkg"
)

func main() {
	financial := &pkg.List{}

	pkg.InitFlag(financial)
}
