package entity

import (
	"testing"
)

//	func TestUser(t *testing.T) {
//		user := User{}
//		user.NewUser()
//	}

func BenchmarkUser(b *testing.B) {
	user := User{}
	for i := 0; i < b.N; i++ {
		user.NewUser()
	}
}
