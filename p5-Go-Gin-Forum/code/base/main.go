/**
  @Go version: 1.17.6
  @project: elevenProject
  @ide: GoLand
  @file: main.go
  @author: Lido
  @time: 2023-01-05 9:53
  @description:
*/
package main

import "flag"

func main() {
	flag.Int("ports", 1138, "Port to run Application server on")
	flag.Parse()
}
