/*
Copyright © 2021 [Company/Organization] Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"github.com/procyon-projects/marker"
	"log"
)

// printErrors prints error(s) if any error exists after processing markers.
func printErrors(errorList marker.ErrorList) {
	if errorList == nil || len(errorList) == 0 {
		return
	}

	for _, err := range errorList {
		switch typedErr := err.(type) {
		case marker.ParserError:
			pos := typedErr.Position
			log.Errorf("%s (%d:%d) : %s\n", typedErr.FileName, pos.Line, pos.Column, typedErr.Error())
		case marker.ErrorList:
			printErrors(typedErr)
		}
	}
}

// validateMarkers visits all files and returns errors
func validateMarkers(collector *marker.Collector, pkgs []*marker.Package) error {
	marker.EachFile(collector, pkgs, func(file *marker.File, fileErrors error) {
		if fileErrors != nil {
			validationErrors = append(validationErrors, fileErrors)
		}
	})

	return marker.NewErrorList(validationErrors)
}