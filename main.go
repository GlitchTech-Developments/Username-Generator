/*
Copyright © 2023 GlitchTech Developments <dev@glitchtech.eu>
*/
package main

import (
	"github.com/GlitchTech-Developments/Username-Generator/cmd"
)

// receive build flags
var versionTag string
var commit string
var buildType string
var versionTagOverride string

// receive build flags
var _versionTag = versionTag
var _commit = commit
var _buildType = buildType
var _versionTagOverride = versionTagOverride

func getVersion(versionTag string, commitHash string, buildType string) string {
	var versionString = ""

	versionString += versionTag + "-" + buildType + "-" + commitHash

	return versionString
}

func main() {
	version := ""
	if _versionTagOverride != "" {
		version += _versionTagOverride
	} else {
		version += getVersion(_versionTag, _commit, _buildType)
	}

	cmd.Execute(version)
}
