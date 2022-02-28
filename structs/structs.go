package structs

type Artists struct {
	Artist []Artist `json:"artists"`
}

type Artist struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}
