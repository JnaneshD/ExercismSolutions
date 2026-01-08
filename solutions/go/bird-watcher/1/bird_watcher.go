package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalBirds := 0
    for i := 0 ; i< len(birdsPerDay); i++ {
        totalBirds = totalBirds + birdsPerDay[i]
    }
    return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weeklen := 7 * (week - 1)
    totalArray := TotalBirdCount(birdsPerDay[weeklen:weeklen+7] )
    return totalArray
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0 ;i < len(birdsPerDay); i++ {
        if (i%2 == 0) {
        	birdsPerDay[i] = birdsPerDay[i] + 1    
        } else {
            birdsPerDay[i] = birdsPerDay[i]
        }
    }
    return birdsPerDay
}
