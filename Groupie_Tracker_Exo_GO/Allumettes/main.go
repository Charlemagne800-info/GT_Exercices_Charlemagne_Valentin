package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Initialisation du générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())

	// Demande du nombre d'allumettes
	var n int
	for n < 4 {
		fmt.Print("Entrez le nombre d'allumettes (minimum 4) : ")
		fmt.Scanln(&n)
	}

	// Détermination du premier joueur
	player := rand.Intn(2) + 1
	fmt.Printf("Le joueur %d commence.\n", player)

	// Boucle principale
	for n > 0 {
		// Affichage du nombre d'allumettes restantes
		fmt.Printf("\n%d allumettes restantes :\n", n)
		for i := 0; i < n; i++ {
			fmt.Print("| ")
		}
		fmt.Println()

		// Demande du nombre d'allumettes à prendre
		var take int
		maxTake := 3
		if n < 3 {
			maxTake = n
		}
		fmt.Printf("Joueur %d, combien d'allumettes voulez-vous prendre (1-%d) ? ", player, maxTake)
		fmt.Scanln(&take)

		// Vérification du nombre d'allumettes à prendre
		if take < 1 || take > maxTake {
			fmt.Println("Nombre d'allumettes invalide.")
			continue
		}

		// Retrait des allumettes prises
		n -= take

		// Changement de joueur
		if player == 1 {
			player = 2
		} else {
			player = 1
		}
	}

	if player == 1 {
		player = 2
	} else {
		player = 1
	}
	// Affichage du résultat
	fmt.Printf("\nJoueur %d a perdu !\n", player)
}
