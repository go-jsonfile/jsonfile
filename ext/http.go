////////////////////////////////////////////////////////////////////////////
// Porgram: http.go
// Purpose: GO HTTP JSON Handling
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: Connor Peet http://stackoverflow.com/questions/17156371/
////////////////////////////////////////////////////////////////////////////

package ext

import (
	"net/http"
	"time"

	"github.com/go-jsonfile/jsonfile"
)

// GetJSON gets JSON from web url.
func GetJSON(url string, js interface{}) error {
	myClient := &http.Client{Timeout: 10 * time.Second}

	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return jsonfile.ReadJSON(r.Body, js)
}
