package helpers

func Contains(s []string, str string) ([]string, []string) {
	var new []string
	var not1 []string
	for _, v := range s {
		if v == str {
			new = append(new, v)
		} else {
			not1 = append(not1, v)
		}
	}

	return new, not1
}
