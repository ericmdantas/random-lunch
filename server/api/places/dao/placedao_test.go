package placedao

import (
	"errors"
	place "github.com/ericmdantas/random-lunch/server/api/places/model"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
)

func fillDB() {
	session, _ := mgo.Dial("localhost")
	defer session.Close()

	for i := 0; i < 10; i++ {
		p := place.Place{Id: bson.NewObjectId(), Name: "some_name", Address: "some_address", Image: "some_image_url", LastTimeChosen: time.Now(), CreatedAt: time.Now()}
		session.DB("random-lunch").C("places").Insert(p)
	}
}

func cleanDB() {
	session, _ := mgo.Dial("localhost")
	defer session.Close()

	session.DB("random-lunch").C("places").RemoveAll(bson.M{})
}

func TestMain(m *testing.M) {
	fillDB()

	m.Run()

	cleanDB()
}

func Test_All(t *testing.T) {
	ps := []place.Place{}

	f := func() {
		All(&ps)
	}

	assert.NotPanics(t, f)
	assert.Equal(t, "some_name", ps[0].Name)
	assert.Equal(t, "some_address", ps[0].Address)
}

func Benchmark_All(b *testing.B) {
	ps := []place.Place{}

	for i := 0; i < b.N; i++ {
		All(&ps)
	}
}

func Test_CreateNew(t *testing.T) {
	p := place.Place{Name: "some_name", Address: "some_address", Image: "some_image_url", LastTimeChosen: time.Now(), CreatedAt: time.Now()}
	pi := place.Place{Name: "", Address: "some_address", Image: "some_image_url", LastTimeChosen: time.Now(), CreatedAt: time.Now()}

	err := CreateNew(&p)

	assert.Nil(t, err)

	err = CreateNew(&p)

	assert.NotEmpty(t, p.Id)

	err = CreateNew(&pi)

	assert.NotEmpty(t, err)
	assert.Equal(t, errors.New("Invalid place."), err)
}
