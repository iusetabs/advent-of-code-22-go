package main

import (
	"testing"
)

func BenchmarkOfImproved(b *testing.B){ 
	main_improved()
}

func BenchmarkOfOld(b *testing.B){ 
	main_old()
}
