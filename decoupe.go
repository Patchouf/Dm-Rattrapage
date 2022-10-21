package main

type Recette struct {
	Bol      []Ingredient
	Maneuvre []Maneuvre
}

type Ingredient struct {
	Nom string
	// en gramme
	Quantite int
}

// Preciser les actions à faire sur les ingredients en fonction de "Type"
type Maneuvre struct {
	// melange | decoupe | cuisson | poubelle
	Type string
	// récupere les ingredients aux index de la liste des ingredients de la recette
	Ingredient []int
}

func Decoupe(Ingredients Ingredient) (Ingredient, Ingredient) {
	var I1 Ingredient
	var I2 Ingredient

	for ing := 0; ing < len(Ingredients.Nom)/2; ing++ {
		I1.Nom += string(Ingredients.Nom[ing])
		I2.Nom += string(Ingredients.Nom[ing+(len(Ingredients.Nom)/2)])
	}
	if len(Ingredients.Nom)%2 != 0 {
		I2.Nom += string(Ingredients.Nom[len(Ingredients.Nom)-1])
	}

	I1.Quantite = Ingredients.Quantite / 2
	I2.Quantite = Ingredients.Quantite / 2
	return I1, I2
}
