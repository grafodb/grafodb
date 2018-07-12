package grafo

type Member struct {
	ID     string `json:"id"`
	Addr   string `json:"addr"`
	Status string `json:"status"`
}

type memberSortByID []Member

func (s memberSortByID) Len() int {
	return len(s)
}
func (s memberSortByID) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s memberSortByID) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}
