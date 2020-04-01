package algorithms

// Les https://en.wikipedia.org/wiki/Bubble_sort
func Bubble_sort_modified(list []int) {
	n := len(list)

	for i := 0; i < n; i++ {
		hold := 0
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j+1]
				list[j+1] = list[j]
				list[j] = temp
				hold = j
			}

		}
		n = hold
	}
}

// Implementering av Bubble_sort algoritmen
func Bubble_sort(list []int) {
	// find the length of list - n
	n := len(list)
	for i := 0; i < n; i++ { //run the following for loop n amount of times, where n = the length of the list of integers

		for j := 0; j < n-1; j++ { //run the following if statement n-1 amount of times
			if list[j] > list[j+1] { //if the item in index j is greater than the next item (j+1),
				temp := list[j+1]   //create a temporary variable that stores the value of that next item (j+1),
				list[j+1] = list[j] //copy the value of j into j+1,
				list[j] = temp      //copy the value temp is holding into index j, effectively swapping the values within
				//increment j by 1, now we will compare what was j+1 and j+2, etcetera
			}

		}
	}
}

// Implementering av Quicksort algoritmen
func QSort(values []int) {
	qsort(values, 0, len(values)-1)
}

func qsort(values []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := values[l]
	i := l + 1

	for j := l; j <= r; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[l], values[i-1] = values[i-1], pivot

	qsort(values, l, i-2)
	qsort(values, i, r)
}
