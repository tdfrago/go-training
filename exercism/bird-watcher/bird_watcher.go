package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, value := range birdsPerDay {
		total += value
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	index := (week - 1) * 7
	return TotalBirdCount(birdsPerDay[index : index+7])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for index, value := range birdsPerDay {
		if (index+1)%2 != 0 {
			birdsPerDay[index] = value + 1
		}
	}
	return birdsPerDay
}
