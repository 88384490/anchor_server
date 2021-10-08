package types

type Movie struct {
	Id            int
	MovieName     string
	MovieType     string
	year          string
	runtime       int
	UserRating    float32
	UserNumber    int
	principalCase string
}

type Anchor struct {
	Id            int
	AnchorName    string
	LastStartTime string
	FanNumber     int
	ChatNumber    int
	Status        bool
}
