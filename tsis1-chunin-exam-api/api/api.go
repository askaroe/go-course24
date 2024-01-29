package api

type Genin struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Clan    string `json:"clan"`
	Village string `json:"village"`
	Team    string `json:"team"`
}

var Genins = []Genin{
	{1, "Naruto", "Uzumaki", "Hidden Leaf", "Team 7"},
	{2, "Sasuke", "Uchiha", "Hidden Leaf", "Team 7"},
	{3, "Gaara", "", "Hidden Sand", "Three Sand Siblings"},
	{4, "Neji", "Hyuga", "Hidden Leaf", "Team Guy"},
	{5, "Temari", "", "Hidden Sand", "Three Sand Siblings"},
	{6, "Kankuro", "", "Hidden Sand", "Three Sand Siblings"},
	{7, "Shikamaru", "Nara", "Hidden Leaf", "Team 10"},
	{8, "Choji", "Akimichi", "Hidden Leaf", "Team 10"},
	{9, "Ino", "Yamanaka", "Hidden Leaf", "Team 10"},
	{10, "Hinata", "Hyuga", "Hidden Leaf", "Team 8"},
}

func GetGeninByName(name string) *Genin {
	for _, genin := range Genins {
		if genin.Name == name {
			return &genin
		}
	}

	return nil
}
