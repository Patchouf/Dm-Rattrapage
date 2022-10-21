func PrintNbr(nombre int) string {
	var Resultat string
	if nombre > 0 && nombre <= 9223372036854775807 {
		new_nombre := nombre
		compteur := 1
		for new_nombre > 9 {
			new_nombre = new_nombre / 10
			compteur *= 10
		}
		new_nombre = nombre
		for compteur+1 > 1 {
			rest := new_nombre / compteur
			new_nombre = new_nombre % compteur
			if rest == 1 {
				Resultat += "1"
			} else if rest == 2 {
				Resultat += "2"
			} else {
				Resultat += string(rune(rest + 48))
			}
			compteur /= 10
		}
	} else if nombre == 0 {
		Resultat += "0"
	} else if nombre < 0 && nombre > -9223372036854775808 {
		nombre = nombre * -1
		Resultat += "-"
		PrintNbr(nombre)
	} else if nombre == -9223372036854775808 {
		nombre = (nombre + 8) / 10 * -1
		Resultat += "-"
		PrintNbr(nombre)
		Resultat += "8"
	}
	return Resultat
}