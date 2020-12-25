package helpers

func StringSliceContains(x string, list []string) bool {
	for _, l := range list {
		if x == l {
			return true
		}
	}
	return false
}
