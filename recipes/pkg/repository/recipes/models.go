package recipes

// Represents a recipe
type Recipe struct {
	Name        string   `json:"name"`
	IsToday     bool     `json:"istoday"`
	Ingredients []string `json:"ingredients"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
}

type Image struct {
	Url string `json:"url"`
}
