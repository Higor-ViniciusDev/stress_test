package entity

type ConcurrencyResult struct {
	TimeStarted int64
	TimeEnded   int64
	Successes   int
	Failures    int
	Codes       map[int]int
}
