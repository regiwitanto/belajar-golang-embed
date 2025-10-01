package belajar_golang_embed

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed version.txt
var version string

//go:embed version.txt
var version2 string

func TestString(t *testing.T) {
	fmt.Println(version)
	fmt.Println(version2)
}