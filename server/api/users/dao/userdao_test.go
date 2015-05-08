package userdao

import (
	"errors"
	user "github.com/ericmdantas/random-lunch/server/api/users/model"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	fillDB()

	m.Run()

	cleanDB()
}

func fillDB() {
	session, _ := mgo.Dial("localhost")
	defer session.Close()

	u := user.User{}

	for i := 0; i < 10; i++ {
		u = user.User{Id: bson.NewObjectId(), Email: "abc123@ajska", Password: "123123", CreatedAt: time.Now()}
		session.DB("random-lunch").C("users").Insert(u)
	}
}

func cleanDB() {
	session, _ := mgo.Dial("localhost")
	defer session.Close()

	session.DB("random-lunch").C("users").RemoveAll(bson.M{})
}

func Test_Consts(t *testing.T) {
	assert.Equal(t, "localhost", URL_DB)
	assert.Equal(t, "random-lunch", DB_NAME)
	assert.Equal(t, "users", COLL)
}

func Test_All(t *testing.T) {
	us := []user.User{}

	All(&us)

	assert.NotEmpty(t, us)
	assert.IsType(t, []user.User{}, us)
	assert.Equal(t, "123123", us[0].Password)
}

func Benchmark_All(b *testing.B) {
	u := []user.User{}

	for i := 0; i < b.N; i++ {

		All(&u)
	}
}

func Test_CreateNew(t *testing.T) {
	u := user.User{Id: bson.NewObjectId(), Email: "ericmdantas0@gmail.com", Password: "123123", CreatedAt: time.Now()}

	f := func() {
		CreateNew(&u)
	}

	assert.NotPanics(t, f)

	ui := user.User{Id: bson.NewObjectId(), Email: "", Password: "", CreatedAt: time.Now()}

	err := CreateNew(&ui)

	assert.Equal(t, errors.New("Invalid user."), err)
}

func Benchmark_CreateNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u := user.User{Id: bson.NewObjectId(), Email: "ericmdantas0@gmail.com", Password: "123123", CreatedAt: time.Now()}
		CreateNew(&u)
	}
}
