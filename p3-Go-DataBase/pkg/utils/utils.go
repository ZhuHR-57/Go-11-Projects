/**
  @Go version: 1.17.6
  @project: elevenProject
  @ide: GoLand
  @file: utils.go
  @author: Lido
  @time: 2023-01-02 16:52
  @description:
*/
package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, X interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), X); err != nil {
			return
		}
	}
}
