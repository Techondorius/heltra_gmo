package dao

import (
	"heltra_gmo/pkg/model"
	"math"
	"time"
)

type RegisterReq struct {
	UserID    string `json:"userid" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Birthdate string `json:"birthdate" binding:"required"`
	Sex       int    `json:"sex" binding:"required"`
	Height    int    `json:"height" binding:"required"`
	Weight    int    `json:"weight" binding:"required"`
}

type LoginReq struct {
	UserID   string `json:"userid" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *RegisterReq) RegisterToUser(user *model.User) {
	user.UserID = r.UserID
	user.Password = r.Password
	user.Name = r.Name
	t, _ := time.Parse(r.Birthdate, "2006-01-02")
	user.Birthdate = t
	user.Sex = r.Sex
	user.Height = r.Height
	user.Weight = r.Weight

	bDTimeTime := time.Unix(user.Birthdate.Unix(), 0)
	age := roundTime((time.Now()).Sub(bDTimeTime).Seconds() / 31207680)

	if true {
		if r.Sex == 1 {
			if age <= 7 {
				user.Objective = 1550
			} else if age <= 9 {
				user.Objective = 1850
			} else if age <= 11 {
				user.Objective = 2250
			} else if age <= 14 {
				user.Objective = 2600
			} else if age <= 17 {
				user.Objective = 2800
			} else if age <= 29 {
				user.Objective = 2650
			} else if age <= 49 {
				user.Objective = 2700
			} else if age <= 64 {
				user.Objective = 2600
			} else if age <= 74 {
				user.Objective = 2400
			} else {
				user.Objective = 2100
			}
		} else {
			if age <= 7 {
				user.Objective = 1450
			} else if age <= 9 {
				user.Objective = 1700
			} else if age <= 11 {
				user.Objective = 2100
			} else if age <= 14 {
				user.Objective = 2400
			} else if age <= 17 {
				user.Objective = 2300
			} else if age <= 29 {
				user.Objective = 2000
			} else if age <= 49 {
				user.Objective = 2050
			} else if age <= 64 {
				user.Objective = 1950
			} else if age <= 74 {
				user.Objective = 1850
			} else {
				user.Objective = 1650
			}
		}
	}
}

func roundTime(input float64) int {
	var result float64

	if input < 0 {
		result = math.Ceil(input - 0.5)
	} else {
		result = math.Floor(input + 0.5)
	}

	// only interested in integer, ignore fractional
	i, _ := math.Modf(result)

	return int(i)
}
