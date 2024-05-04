package dkn

import (
	"fmt"
	"github.com/killrpenguin/dkn_tool"
)


func main() {
	fmt.Println("Hello World!")
	resp := DnRequest("https://raw.githubusercontent.com/a8m/golang-cheat-sheet/9c634033efeb0fdd45c8b08030be9d53d3c4ea73/README.md")
	fmt.Println(resp)
}
