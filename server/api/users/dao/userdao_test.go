package userdao

import (
	_ user "github.com/ericmdantas/random-lunch/server/api/users//model"
	_ "github.com/stretchr/testify/assert"
	_ "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	_ "testing"
	_ "time"
)
