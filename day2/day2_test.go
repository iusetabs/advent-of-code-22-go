package main

import (
	"testing"
)

func BenchmarkOfMaps(b *testing.B){ 
	main_improved()
}

func BenchmarkOfIfs(b *testing.B){ 
	main_old()
}
