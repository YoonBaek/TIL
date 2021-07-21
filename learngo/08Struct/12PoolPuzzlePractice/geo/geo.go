package geo

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// Embedding Practice.
type Landmark struct {
	Name string
	Coordinates
}
