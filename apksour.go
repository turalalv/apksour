/*
 made with love by @Legendhot 9/2020
*/

package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/turalalv/apksour/command/apktool"
	dependency "github.com/turalalv/apksour/command/dependency"
	"github.com/turalalv/apksour/directory"
	"github.com/turalalv/apksour/extractor"
	"os"
)



func main() {

	parser := argparse.NewParser("apksour", "ApkSouR - Extract endpoints from APK files")
	apk := parser.String("a", "apk", &argparse.Options{Required: true, Help: "Input a path to APK file."})

	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(-1)
	}

	var baseApk = *apk
	var tempDir = directory.CreateTempDir()

	dependency.AreAllReady()
	apktool.RunApktool(baseApk, tempDir)
	extractor.Extract(tempDir)
	directory.RemoveTempDir(tempDir)
}
