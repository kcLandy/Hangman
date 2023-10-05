package main

func join(strings []string, separator string) string {
	if len(strings) == 0 {
		return ""
	}
	var s string = ""
	var lastIdx int = len(strings) - 1
	for _, v := range strings[0:lastIdx] {
		s += string(v) + separator
	}
	return s + strings[lastIdx]
}