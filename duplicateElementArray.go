package tasks	

func HasDuplicate (numbers []int) string {

	var duplicate bool

	for i, number := range numbers {

		for j := i; j <len(numbers) - 1; j++ {

			if number == numbers[j+1] {
				duplicate = true
				break
			}
		}

	}
	if duplicate {
		return "Array has duplicate element"
	} else {
		return "Array has not duplicate element"
	}
}