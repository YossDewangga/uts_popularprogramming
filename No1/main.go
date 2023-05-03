// Elipsis
package main

import (
	"fmt"
)

func total(items ...int) int {
	jumlah := 0
	for _, item := range items {
		jumlah += item
	}
	return jumlah
}

func main() {
	ayamgoreng := 15000
	ayambakar := 18000

	jml_item_1 := 2
	jml_item_2 := 1

	Bayar := total(ayamgoreng*jml_item_1, ayambakar*jml_item_2)
	fmt.Printf("Total: %d \n", Bayar)
}
