package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools

type Rank [8]bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	occupiedCount := 0
	for _, v := range cb[rank] {
		if v {
			occupiedCount++
		}
	}

	return occupiedCount
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	occupiedCount := 0
	index := file - 1

	if 7 < index {
		return occupiedCount
	}

	for _, v := range cb {
		if v[index] {
			occupiedCount++
		}
	}

	return occupiedCount
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	count := 0
	for _, v := range cb {
		for range v {
			count++
		}
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, v := range cb {
		for _, r := range v {
			if r {
				count++
			}
		}
	}

	return count
}
