package main

import (
	"fmt"
	//Add these imports here too so that they get go gotten
	_ "bitbucket.org/creachadair/cityhash"
	_ "bitbucket.org/jpathy/dmc/cityhash"
	_ "github.com/dgryski/go-spooky"
	_ "github.com/huichen/murmur"
	_ "github.com/leemcloughlin/gofarmhash"
	_ "github.com/pborman/uuid"
	_ "github.com/reusee/mmh3"
	_ "github.com/tildeleb/hashland/spooky"
	_ "github.com/zhangxinngang/murmur"
)

func main() {
	fmt.Println("Please run this with 'go test -bench .'")
}
