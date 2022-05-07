package domain

import "time"

type User struct {
	Id             	string
	Firstname      	string
	Email          	string
	MobileNumber   	string
	Gender         	string
	BirthDay       	time.Time
	Username       	string
	Biography      	string
	Experience     	string
	Education      	string
	Skills         	string
	Interests      	string
	Password       	string
	FollowingUsers 	[]string
	FollowedByUsers []string
	Public       	bool
}

type Session struct {
	Id     string
	UserId string
	Date   time.Time
	Role   string
}
