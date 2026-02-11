package api

type Locations struct {
	Next     string
	Previous string
	Results  []Location
}

type Location struct {
	Name string
}

type LocationArea struct {
	Pokemons []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name       string
	Height     int `json:"height"`
	Weight     int `json:"weight"`
	Experience int `json:"base_experience"`

	Stats []struct {
		Value int `json:"base_stat"`

		StatInfo struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`

	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}
