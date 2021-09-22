interface RecipeInput {
	ID:             number;
	Title:          string;
	Description:    string;
    Ingredients: Array<IngredientInput>;
	Method:         string;
	PrepTime:       number;
	CookTime:       number;
	TotalTime:      number;
	Amount_Persons: number;
}

interface IngredientInput {
    ID: number;
    amount: number;
    unit: string;
}