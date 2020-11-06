package main

import "testing"

type DOBTests struct {
	Dob string
	AgeInYears int
	Err error
}

func TestGetTime(t *testing.T) {
	testTable := []DOBTests{
		{
			Dob: "2020-01-01",
			AgeInYears: 0,
			Err: nil,
		},
		{
			Dob: "1995-09-20",
			AgeInYears: 25,
			Err: nil,
		},
	}

	for _, table := range testTable {
		ageInYears, err := GetTime(table.Dob)
		if !(ageInYears == table.AgeInYears && err == table.Err) {
			t.Errorf("Incorrect output. AgeInYears: %d, want: %d", ageInYears, table.AgeInYears)

		}
	}
	//ageInYears, err := GetTime("2010-01-01")
	//if !(ageInYears == 10 && err == nil) {
	//	t.Errorf("Incorrect output. AgeInYears: %d, want: %d", ageInYears, 10)
	//}
}