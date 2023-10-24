/*
Copyright © 2023 GlitchTech Developments <dev@glitchtech.eu>
*/
package main

import (
	"embed"

	"github.com/GlitchTech-Developments/Username-Generator/cmd"
)

//go:embed current-version.txt
var f embed.FS

// type Config struct {
// 	version string
// }

func main() {
	version, _ := f.ReadFile("current-version.txt")
	cmd.Execute(string(version))
}
