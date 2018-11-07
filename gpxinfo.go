// Copyright 2013, 2014 Peter Vasil, Tomo Krajina. All
// rights reserved. Use of this source code is governed
// by a BSD-style license that can be found in the
// LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/mbecker/gpxgo/gpx"
)

func main() {
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	gpxDirectory := filepath.Join(currentDirectory, "gpx_files")
	files, readDirErr := ioutil.ReadDir(gpxDirectory)
	if readDirErr != nil {
		panic(readDirErr)
	}
	for _, file := range files {
		gpxFile, gpxParseFileErr := gpx.ParseFile(filepath.Join(gpxDirectory, file.Name()))
		if gpxParseFileErr != nil {
			fmt.Printf("%s error parsing file: %s", file.Name(), gpxParseFileErr)
			continue
		}
		fmt.Println(gpxFile.GetGpxInfo())
	}
}
