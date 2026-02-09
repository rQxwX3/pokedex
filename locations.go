package main

type Locations struct {
	Count    int
	Next     string
	Previous string
	Results  []Location
}

type Location struct {
	Name string
	Url  string
}
