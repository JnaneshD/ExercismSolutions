package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool
// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File
// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0 
    for k, v := range(cb) {
        if k == file {
            for _, i := range(v) {
                if i == true {
                    count ++ 
                }
            }
        }
    }
    return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    if rank < 1 || rank > 8 {
		return 0
	}
	total_occupied_rank := 0
    for _, v := range(cb ) {
        if rank-1 < len(v) && v[rank-1] {
            total_occupied_rank++
        }
    }
    return total_occupied_rank
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	// Now count the squares, it is just two for loops O(n2)
    total_squares := 0
    for _,v := range(cb) {
        for range(v) {
            total_squares++
        }
    }
    return total_squares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	total_occupied := 0 
    for _,v := range(cb) {
        for _, val := range(v) {
            if val == true{
                total_occupied++
            }
        }
    }
    return total_occupied
}
