package main

func Find(array []string, values ...string) (found bool) {
	for _, arrayItem := range array {
		for _, value := range values {
			if arrayItem == value {
				found = true
				return
			}
		}
	}

	return
}
