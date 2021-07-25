package main

import (
	"github.com/diazharizky/codebase-rest-golang/cmd"
	"github.com/diazharizky/codebase-rest-golang/configs"
)

func main() {
	configs.LoadConfig()
	cmd.Execute()
}
