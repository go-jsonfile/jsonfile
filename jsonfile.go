////////////////////////////////////////////////////////////////////////////
// Porgram: jsonfile.go
// Purpose: GO JSON File Handling
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: Mkideal Wang https://github.com/mkideal/cli/issues/22
////////////////////////////////////////////////////////////////////////////

package jsonfile

import (
	"encoding/json"
	"io"
	"os"
)

// ReadJSON reads JSON from stream.
func ReadJSON(r io.Reader, js interface{}) error {
	return json.NewDecoder(r).Decode(js)
}

// ReadJSONFromFile reads JSON from the file with given name.
func ReadJSONFromFile(filename string, js interface{}) error {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return err
	}
	return ReadJSON(file, js)
}

// WriteJSON writes JSON to the stream.
func WriteJSON(w io.Writer, js interface{}) error {
	return json.NewEncoder(w).Encode(js)
}

// WriteJSONFromFile writes JSON to the file with given name.
func WriteJSONToFile(filename string, js interface{}) error {
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return err
	}
	return WriteJSON(file, js)
}
