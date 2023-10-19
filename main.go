/*
Copyright © 2023 GlitchTech Developments <dev@glitchtech.eu>
*/
package main

import (
	"github.com/GlitchTech-Developments/Username-Generator/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.EnableCaseInsensitive = true
	cmd.Execute()
}
