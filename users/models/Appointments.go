package models

import "time"

type Appointment struct {
	Id             uint32    `json:"id"`
	Doctor         User      `json:"doctor"`
	Patient        User      `json:"patient"`
	Speciality     uint32    `json:"speciality"`
	Sub_speciality uint32    `json:"sub_speciality"`
	Details        string    `json:"details"`
	Duration       uint32    `json:"duration"`
	Cost           float32   `json:"cost"`
	Start_at       time.Time `json:"start_at"`
	End_at         time.Time `json:"end_at"`
	Status         string    `json:"status"`
	Created_at     time.Time `json:"created_at"`
	Updated_at     time.Time `json:"updated_at"`
}
