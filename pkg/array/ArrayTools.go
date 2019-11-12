package array


func StringArrayEquals(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for _, va := range a {
		found := false
		for _, vb := range b {
			if va == vb {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func StringArrayContains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
