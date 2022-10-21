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

func Melange(Ingredients []Ingredient) Ingredient {
	Ingredient3 := Ingredient{
		Nom:      "a",
		Quantite: 4,
	}

	Ingredient3.Nom = (Ingredients[0].Nom + Ingredients[1].Nom)
	Ingredient3.Quantite = (Ingredients[0].Quantite + Ingredients[1].Quantite)

	return Ingredient3

}
