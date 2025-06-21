package main

func generateId() func() int {
	var count int
	return func() int {
		count++
		return count
	}
}
