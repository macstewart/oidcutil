package main

import (
	"embed"
	"ssorry/cmd"
	_ "ssorry/internal/store"
)

//go:embed resources/html
var resources embed.FS

func main() {
	cmd.StoreResources(resources)
	cmd.Execute()
}
