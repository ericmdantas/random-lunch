package usermodel

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
)

var testUsers = []struct {
	in  User
	out bool
}{
	{
		in:  User{Id: bson.NewObjectId(), Email: "", Password: "123223", CreatedAt: time.Now()},
		out: false,
	},
	{
		in:  User{Id: bson.NewObjectId(), Email: "a", Password: "123223", CreatedAt: time.Now()},
		out: false,
	},
	{
		in:  User{Id: bson.NewObjectId(), Email: "ericdantas0@gmail.com", Password: "123223", CreatedAt: time.Now()},
		out: true,
	},
}

func Test_IsValid(t *testing.T) {
	for _, tu := range testUsers {
		b := tu.in.IsValid()

		assert.Equal(t, tu.out, b)
	}
}

func Benchmark_IsValid(b *testing.B) {

	u := User{Id: bson.NewObjectId(), Email: "ericdantas0@gmail.com", Password: "123223", CreatedAt: time.Now()}

	for i := 0; i < b.N; i++ {
		u.IsValid()
	}
}
