package suppliers 

type Supplier struct {
	Name        string       `json:"name"`
  Description string       `json:"description"`
  IsFavorite  bool         `json:"isfavorite"`
  Products []Product
  Image       string       `json:"image"`
}

type Product struct {
  Name string `json:"name"`
  Description string `json:"description"`
  Image string `json:"image"`
}

