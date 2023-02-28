package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Initialisation du générateur de nombres aléatoires

	var n int // Demande du nombre d'allumettes
	for n < 4 {
		fmt.Print("Entrez le nombre d'allumettes (minimum 4) : ")
		fmt.Scanln(&n)
	}

	player := rand.Intn(2) + 1 // Détermination du premier joueur
	fmt.Printf("Le joueur %d commence.\n", player)

	// Boucle principale
	for n > 0 {
		fmt.Printf("\n%d allumettes restantes :\n", n) // Affichage du nombre d'allumettes restantes
		for i := 0; i < n; i++ {
			fmt.Print("| ") //Affichage graphique des allumettes
		}
		fmt.Println()

		var take int // Demande du nombre d'allumettes à prendre
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
