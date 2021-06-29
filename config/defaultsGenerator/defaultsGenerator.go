// Copyright (C) 2019-2021 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/algorand/go-algorand/config"
)

const defaultLocalVariableDeclaration = "var defaultLocal"

var outputfilename = flag.String("o", "", "Name of the file where the generated datastructure would go to.")
var packageName = flag.String("p", "", "Name of the package.")
var headerFileName = flag.String("h", "", "Name of the header filename")
var jsonExampleFileName = flag.String("j", "", "Name of the json example file")

var autoGenHeader = `
// This file was auto generated by ./config/defaultsGenerator/defaultsGenerator.go, and SHOULD NOT BE MODIFIED in any way
// If you want to make changes to this file, make the corresponding changes to Local in config.go and run "go generate".
`

// printExit prints the given formatted string ( i.e. just like fmt.Printf ), with the defaultGenerator executable program name
// at the beginning, and exit the process with a error code of 1.
func printExit(fmtStr string, args ...interface{}) {
	fmt.Printf("%s: "+fmtStr, append([]interface{}{filepath.Base(os.Args[0])}, args...)...)
	os.Exit(1)
}

func main() {

	flag.Parse()
	if *outputfilename == "" || *packageName == "" || *headerFileName == "" || *jsonExampleFileName == "" {
		printExit("one or more of the required input arguments was not provided\n")
	}

	localDefaultsBytes, err := ioutil.ReadFile(*headerFileName)
	if err != nil {
		printExit("Unable to load file %s : %v", *headerFileName, err)
	}
	localDefaultsBytes = []byte(strings.Replace(string(localDefaultsBytes), "{DATE_Y}", fmt.Sprintf("%d", time.Now().Year()), 1))
	localDefaultsBytes = append(localDefaultsBytes, []byte(autoGenHeader)...)

	// add the package name:
	localDefaultsBytes = append(localDefaultsBytes, []byte(fmt.Sprintf("\npackage %s\n\n", *packageName))...)

	autoDefaultsBytes := []byte(prettyPrint(config.AutogenLocal, "go"))

	localDefaultsBytes = append(localDefaultsBytes, autoDefaultsBytes...)

	err = ioutil.WriteFile(*outputfilename, localDefaultsBytes, 0644)
	if err != nil {
		printExit("Unable to write file %s : %v", *outputfilename, err)
	}

	// generate an update json for the example as well.
	autoDefaultsBytes = []byte(prettyPrint(config.AutogenLocal, "json"))
	err = ioutil.WriteFile(*jsonExampleFileName, autoDefaultsBytes, 0644)
	if err != nil {
		printExit("Unable to write file %s : %v", *jsonExampleFileName, err)
	}
}

type byFieldName []reflect.StructField

func (a byFieldName) Len() int      { return len(a) }
func (a byFieldName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byFieldName) Less(i, j int) bool {
	if a[i].Name == "Version" {
		return true
	} else if a[j].Name == "Version" {
		return false
	}
	return a[i].Name < a[j].Name
}

func prettyPrint(c config.Local, format string) (out string) {
	localType := reflect.TypeOf(c)

	fields := []reflect.StructField{}
	for fieldNum := 0; fieldNum < localType.NumField(); fieldNum++ {
		fields = append(fields, localType.Field(fieldNum))
	}

	sort.Sort(byFieldName(fields))

	if format == "go" {
		out = fmt.Sprintf("%s = Local{\n", defaultLocalVariableDeclaration)
	} else {
		out = "{\n"
	}

	for fieldIdx, field := range fields {
		switch field.Type.Kind() {
		case reflect.Bool:
			v := reflect.ValueOf(&c).Elem().FieldByName(field.Name).Bool()
			if format == "go" {
				out = fmt.Sprintf("%s\t%s:\t%v,\n", out, field.Name, v)
			} else {
				out = fmt.Sprintf("%s    \"%s\": %v,\n", out, field.Name, v)
			}
		case reflect.Int32:
			fallthrough
		case reflect.Int:
			fallthrough
		case reflect.Int64:
			v := reflect.ValueOf(&c).Elem().FieldByName(field.Name).Int()
			if format == "go" {
				out = fmt.Sprintf("%s\t%s:\t%d,\n", out, field.Name, v)
			} else {
				out = fmt.Sprintf("%s    \"%s\": %d,\n", out, field.Name, v)
			}
		case reflect.Uint32:
			fallthrough
		case reflect.Uint:
			fallthrough
		case reflect.Uint64:
			v := reflect.ValueOf(&c).Elem().FieldByName(field.Name).Uint()
			if format == "go" {
				out = fmt.Sprintf("%s\t%s:\t%d,\n", out, field.Name, v)
			} else {
				out = fmt.Sprintf("%s    \"%s\": %d,\n", out, field.Name, v)
			}
		case reflect.String:
			v := reflect.ValueOf(&c).Elem().FieldByName(field.Name).String()
			if format == "go" {
				out = fmt.Sprintf("%s\t%s:\t\"%s\",\n", out, field.Name, v)
			} else {
				out = fmt.Sprintf("%s    \"%s\": \"%s\",\n", out, field.Name, v)
			}
		case reflect.Map:
			if reflect.ValueOf(&c).Elem().FieldByName(field.Name).Len() == 0 {
				if format == "go" {
					// it's an empty map; good, we know how to initialize empty maps.
					mapKeysType := field.Type.Key()
					mapValueType := field.Type.Elem()

					out = fmt.Sprintf("%s\t%s:\tmap[%s]%s{},\n", out, field.Name, mapKeysType, mapValueType)
				} else {
					out = fmt.Sprintf("%s    \"%s\": {},\n", out, field.Name)
				}
			} else {
				printExit("non-empty default maps data type encountered when reflecting on config.Local datatype %s", field.Name)
			}
		default:
			printExit("unsupported data type (%s) encountered when reflecting on config.Local datatype %s", field.Type.Kind(), field.Name)
		}
		if format != "go" {
			if fieldIdx == len(fields)-1 {
				out = out[:len(out)-2] + "\n"
			}
		}
	}
	if format == "go" {
		out = out + "}"
	} else {
		out = out + "}\n"
	}
	return
}
