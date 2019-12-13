package main

type stringSlice []string

func (p stringSlice) Len() int {
	return len(p)
}

func (p stringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p stringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	TestHttp1()
}
