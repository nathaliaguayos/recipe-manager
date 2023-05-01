package meals

type Meal struct {
	ID          string                 `json:"ID"`
	Name        string                 `json:"nombre"`
	Label       string                 `json:"tipo"`
	Ingredients []interface{}          `json:"ingredientes"`
	Procedure   []interface{}          `json:"preparacion"`
	Image       string                 `json:"imagen"`
	Kcal        string                 `json:"kcal"`
	Macros      map[string]interface{} `json:"macros"`
}

type Macros struct {
	Lipids  string `json:"lipidos"`
	Carbs   string `json:"carbohidratos"`
	Protein string `json:"proteina"`
}

// TODO: Meal struct must be like the following: I need to found the proper way to deserealize the JSON object.
type TestMeal struct {
	ID          string   `json:"ID"`
	Name        string   `json:"nombre"`
	Label       string   `json:"tipo"`
	Ingredients []string `json:"ingredientes"`
	Procedure   []string `json:"preparacion"`
	Image       string   `json:"imagen"`
	Kcal        string   `json:"kcal"`
	Macros      Macros   `json:"macros"` //TODO: I need to see how I can handle macro's items with a map[string]interface{}
}
