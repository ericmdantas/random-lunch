package placemodel

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
)

var testPlacesIsValid = []struct {
	in  Place
	out bool
}{
	{
		in:  Place{Id: bson.NewObjectId(), Name: "", Address: "", Image: "somewhere.jpg", LastTimeChosen: time.Now(), CreatedAt: time.Now()},
		out: false,
	},
	{
		in:  Place{Id: bson.NewObjectId(), Name: "", Address: "a", Image: "somewhere.jpg", LastTimeChosen: time.Now(), CreatedAt: time.Now()},
		out: false,
	},
	{
		in:  Place{Id: bson.NewObjectId(), Name: "a", Address: "a", Image: "somewhere.jpg", LastTimeChosen: time.Now(), CreatedAt: time.Now()},
		out: true,
	},
}

var testsPlacesIsNew = []struct {
	in  Place
	out bool
}{
	{
		in:  Place{Id: bson.NewObjectId(), Name: "", Address: "", Image: "somewhere.jpg", LastTimeChosen: time.Now(), CreatedAt: time.Now()},
		out: false,
	},
	{
		in:  Place{Id: "", Name: "", Address: "", Image: "somewhere.jpg", LastTimeChosen: time.Now(), CreatedAt: time.Now()},
		out: true,
	},
}

func Test_IsValid(t *testing.T) {
	for _, tp := range testPlacesIsValid {
		r := tp.in.IsValid()

		assert.Equal(t, tp.out, r)
	}
}

func Benchmark_IsValid(b *testing.B) {
	p := Place{Id: bson.NewObjectId(), Name: "a", Address: "a", Image: "somewhere.jpg", LastTimeChosen: time.Now(), CreatedAt: time.Now()}

	for i := 0; i < b.N; i++ {
		p.IsValid()
	}
}

func Test_IsNew(t *testing.T) {
	for _, tp := range testsPlacesIsNew {
		act := tp.in.IsNew()

		assert.Equal(t, tp.out, act)
	}
}

func Benchmark_IsNew(b *testing.B) {
	p := Place{Id: bson.NewObjectId(), Name: "a", Address: "a", Image: "somewhere.jpg", LastTimeChosen: time.Now(), CreatedAt: time.Now()}

	for i := 0; i < b.N; i++ {
		p.IsNew()
	}
}
