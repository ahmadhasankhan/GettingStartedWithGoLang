package main

type Output struct {
	Description string  `json:"description"`
	Cereal      float32 `json:"cereal"`
	Milk        float32 `json:"milk"`
}

type Recommendations []Output
