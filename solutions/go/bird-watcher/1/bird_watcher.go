package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	birdAmountTotal := 0
	for _, count := range birdsPerDay {
		birdAmountTotal += count
	}
	return birdAmountTotal
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	startIndex := (week - 1) * 7
	endIndex := startIndex + 7
	birdAmountTotal := 0

	for _, count := range birdsPerDay[startIndex:endIndex] {
		birdAmountTotal += count
	}

	return birdAmountTotal
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for index, count := range birdsPerDay {
		if index%2 == 0 {
			birdsPerDay[index] = count + 1
		}
	}

	return birdsPerDay
}
