package main

import (
	"encoding/hex"
	"fmt"

	"github.com/devedge/imagehash"
	"github.com/steakknife/hamming"
)

func main() {
	src512, _ := imagehash.OpenImg("../testdata/a.jpg")
	src256, _ := imagehash.OpenImg("../testdata/b.jpg")
	srcInv, _ := imagehash.OpenImg("../testdata/a.jpg")

	hash512, _ := imagehash.Dhash(src512, 8)
	hash256, _ := imagehash.Dhash(src256, 8)
	hashInv, _ := imagehash.Dhash(srcInv, 8)

	// Hamming distance of 0, since the images are simply different sizes
	fmt.Println("'lena_512.png' dhash:", hex.EncodeToString(hash512))
	fmt.Println("'lena_256.png' dhash:", hex.EncodeToString(hash256))
	fmt.Println("The Hamming distance between these:", hamming.Bytes(hash512, hash256))
	fmt.Println()

	// Completely different dhash, since an inverted image has a completely
	// different gradient colorscheme
	fmt.Println("'lena_512.png' dhash:         ", hex.EncodeToString(hash512))
	fmt.Println("'lena_inverted_512.png' dhash:", hex.EncodeToString(hashInv))
	fmt.Println("The Hamming distance between these:", hamming.Bytes(hash512, hashInv))

}
