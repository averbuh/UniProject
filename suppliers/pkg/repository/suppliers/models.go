package suppliers

type Supplier struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsFavorite  bool   `json:"isfavorite"`
	Image       string `json:"image"`
	// Products []Product
}

type Image struct {
	Url string `json:"url"`
}

// type Product struct {
//   Name string `json:"name"`
//   Description string `json:"description"`
//   Image string `json:"image"`
// }
