/* vi: set ft=go ts=4 sw=4 sts=0 noet: */

package main

import (
	"fmt"
	"github.com/moby/moby/pkg/namesgenerator"
	"math/rand"
	"strings"
	"time"
)

// GetRandomNameWithHyphen generates a random name with hyphen. If retry is non-zero, a random
// integer between 0 and 10 will be added to the end of the name, e.g `focused-turing3`
func GetRandomNameWithHyphen(retry int) string {
	name := strings.Replace(namesgenerator.GetRandomName(retry), "_", "-", 1)
	return name
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(GetRandomNameWithHyphen(0))
}

