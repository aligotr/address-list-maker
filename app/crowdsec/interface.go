package crowdsec

type addressesData struct {
	address  string
	scenario string
}

type origin map[string][]addressesData

type cache struct {
	v4 origin
	v6 origin
}
