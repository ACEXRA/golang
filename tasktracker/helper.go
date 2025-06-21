package main


func id() func() int {
	var count int;
	return func() int {
		 count++
		 return count
	}
}
