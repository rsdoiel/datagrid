//
// datagrid.go is a Go package working with 2d and 3d datastructures like CSV files, xlsx Workbooks or Google Sheets.
// datagrid provides some common abstractions for moving data between formats as well as processing 2d grid oriented
// data.
//
// @author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2017, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package datagrid

import (
	"encoding/json"
	"fmt"
)

// Cell holds the smallest unit of a grid (a 2d array of cells)
type Cell struct {
	// Value holds the value associated with the cell, typically
	// it corresponds to a Go type
	Value interface{}
	// Format holds a format string suitable for use with functions like fmt.Sprintf() and log.Printf()
	Format string
}

type Grid [][]Cell

type Sheet interface {
	Connection(map[string]string) error
	ReadSheet(string, string, string) (*Grid, error)
	WriteSheet(string, string, string, *Grid) error
}

type Workbook struct {
	Sheets []Sheet
}

// Address holds the x/y location of a cell in a grid. "0,0" is the top left and "n,n" would be the lower right
type Address [2]int

// CellRange takes a string like A1:A16 and returns a list of cell addresses identitified as belonging to
// the range.
//
// return an array of Address or an error
//
func CellRange(s string) ([]Address, error) {
	addresses := []Address{}

	// FIXME: need something that can turn a column letter into an appropriate column index number
	return addresses, fmt.Errorf("CellRange() not implented")
}

func (c *Cell) String() string {
	fstring := c.Format
	if formatString == "" {
		switch c.Value.(type) {
		case int:
			formatString = "%d"
		case int64:
			formatString = "%d"
		case float64:
			formatString = "%d"
		case time.Time:
			formatString = "%s"
		case string:
			formatString = "%s"
		case map[string]interface{}:
			if src, err := json.Marshal(c.Value); err != nil {
				return ""
			} else {
				return fmt.Sprintf("%s", src)
			}
		case []interface{}:
			if src, err := json.Marshal(c.Value); err != nil {
				return ""
			} else {
				return fmt.Sprintf("%s", src)
			}
		case json.Number:
			return json.Number.String()
		default:
		}
	}
	return fmt.Sprintf(formatString, c.Value)
}
