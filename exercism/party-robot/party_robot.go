package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	output := "Welcome to my party, " + name + "!"
	return output
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	age_string := fmt.Sprint(age)
	output := "Happy birthday " + name + "! You are now " + age_string + " years old!"
	return output
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	table_string := fmt.Sprintf("%03v", table)
	distance_string := fmt.Sprintf("%.1f", distance)
	output := "Welcome to my party, " + name + "!\nYou have been assigned to table " + table_string + ". Your table is " + direction + ", exactly " + distance_string + " meters from here.\nYou will be sitting next to " + neighbor + "."
	return output
}
