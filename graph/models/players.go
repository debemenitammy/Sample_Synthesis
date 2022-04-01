package models

// Struct for `players`
type Players struct {
	playerid int32
	cohortid int32
	name string
	email string
	age int32
	dob string
	username string
	password string
	admin int8
	order int8
	reset_token string
	password_sent string
	last_login string
	timezone string
	at_id string
	hashed_password string
	start_date string
	user_id string
}