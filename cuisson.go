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

// PS J'AI LA FLEMME DE CRéé UN ANNEXE

// Preciser les actions à faire sur les ingredients en fonction de "Type"
type Maneuvre struct {
	// melange | decoupe | cuisson | poubelle
	Type string
	// récupere les ingredients aux index de la liste des ingredients de la recette
	Ingredient []int
}

func Cuisson(Ingr Ingredient) Ingredient {
	var E Ingredient

	for ing := len(Ingr.Nom) - 1; ing >= 0; ing-- {
		E.Nom += string(Ingr.Nom[ing])
	}
	E.Quantite = Ingr.Quantite * 9 / 10

	return E
}
