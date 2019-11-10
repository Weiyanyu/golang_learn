package main

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i += 1
		}
	}
	return strings[:i]
}
