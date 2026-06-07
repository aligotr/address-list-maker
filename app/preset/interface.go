package preset

type platform struct {
	extension string
}

type protocol struct {
	Name string
	Sep  string
}

type protocols struct {
	V4 protocol
	V6 protocol
}
