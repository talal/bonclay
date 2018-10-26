/**********************************************************************************
*
* Copyright 2018 Muhammad Talal Anwar <talalanwar@protonmail.com>
*
* This program is free software: you can redistribute it and/or modify it under
* the terms of the GNU General Public License as published by the Free Software
* Foundation, either version 3 of the License, or (at your option) any later
* version.
*
* This program is distributed in the hope that it will be useful, but WITHOUT ANY
* WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
* A PARTICULAR PURPOSE. See the GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License along with
* this program. If not, see <http://www.gnu.org/licenses/>.
*
**********************************************************************************/

package mistertwo

import (
	"fmt"

	"github.com/alecthomas/kingpin"
)

const (
	dstProblem     = "problem with destination"
	srcProblem     = "problem with source"
	srcPDirProblem = "problem with source's parent directory"
	targetProblem  = "problem with target"

	dCopyFailure   = "could not copy directory"
	fCopyFailure   = "could not copy file"
	dCleanFailure  = "could not clean directory"
	lCreateFailure = "could not remove broken symlink"
	lRemoveFailure = "could not remove broken symlink"
	syncFailure    = "could not sync"

	dstExists       = "destination already exists"
	srcNotExists    = "source does not exist"
	targetNotExists = "target does not exist"

	dirSkip     = "directory skipped"
	fileSkip    = "file skipped"
	linkSkip    = "symlink skipped"
	unknownType = "unknown type: neither file nor directory"
)

// newError is a helper function that for a given string and value returns a
// new error in the form of "string: value"
func newError(str string, val interface{}) error {
	return fmt.Errorf("%s: %v", str, val)
}

func printTaskErrors(taskName string, errors []string) {
	var uniqueErrors []string
	var errExists = make(map[string]bool)
	for _, v := range errors {
		if _, exists := errExists[v]; !exists {
			uniqueErrors = append(uniqueErrors, v)
			errExists[v] = true
		}
	}

	fmt.Println(withColorf(red, "\nSome errors occured during %s:", taskName))
	for _, v := range uniqueErrors {
		fmt.Printf("\t%s\n", v)
	}
}

// fatalIfError is a convenient wrapper for kingpin.FatalIfError().
func fatalIfError(err error, str string) {
	kingpin.FatalIfError(err, str)
}
