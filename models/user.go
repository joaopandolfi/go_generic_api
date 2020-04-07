package models

import "gopkg.in/mgo.v2/bson"

const USER_ROOT int = 1
const USER_ADMIN int = 2
const USER_ESPECIALIST int = 3
const USER_DOCTOR int = 4
const USER_SECRETARY int = 5
const USER_CLIENT int = 6

const SESSION_VALUE_INSTITUTION string = "institution"
const SESSION_VALUE_TOKEN string = "token"
const SESSION_VALUE_LEVEL string = "level"
const SESSION_VALUE_USERNAME string = "username"
const SESSION_VALUE_ID string = "user_id"
const SESSION_VALUE_NAME string = "name"
const SESSION_VALUE_LOGGED string = "logged"
const SESSION_VALUE_SPECIALTY string = "specialty"

type User struct {
	People
	Uid       bson.ObjectId `bson:"_id,omitempty"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Token     string `json:"token"`
	Picture   string `json:"foto"`
	Password  string `json:"password"`
	ID        int  	`json:"iduser"`
	Level     int    `json:"level"`
	Instution int    `json:"idcompany"`
	Specialty int
}
