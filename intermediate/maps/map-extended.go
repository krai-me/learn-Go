package main

import "fmt"

func main() {
	// Initialize a map with Pokémon names as keys and their types as values
	pokemons := map[string]string{
		"Pikachu":    "Electric",
		"Charmander": "Fire",
		"Bulbasaur":  "Grass/Poison",
		"Squirtle":   "Water",
		"Gengar":     "Ghost/Poiso",
	}

	// Try to access the type of a Pokémon that is not in the map ("Eevee")
	// 'pokemon' will store the value associated with the key (if found), which is the Pokémon's type.
	// 'status' will be a boolean indicating whether the key "Eevee" exists in the map.
	pokemon, status := pokemons["Eevee"]

	// Check if the key "Eevee" was found in the map
	if status {
		// If found, print the Pokémon's type
		fmt.Println("We found the type:", pokemon)
	} else {
		// If not found, indicate the Pokémon is not in the map
		fmt.Println("Pokémon not found")
	}

	// Adding and Modifying:
	// To add or modify an entry in the map, we use the same syntax.
	// If the key is already present, it modifies the existing value.
	// If the key is not present, it adds a new key-value pair to the map.
	// Syntax: yourMap[newKey] = newValue

	pokemons["Jigglypuff"] = "Normal/Fairy" //adds new pokemon as key and value
	pokemons["Gengar"] = "Ghost/Poison"     //updates the value of genegar key
	fmt.Println(pokemons)                   //map[Bulbasaur:Grass/Poison Charmander:Fire Gengar:Ghost/Poison Jigglypuff:Normal/Fairy Pikachu:Electric Squirtle:Water]

	// Removing
	// Syntax: delete(yourMap,keyValueToDelete)
	delete(pokemons, "Gengar") //removes gengar key
	fmt.Println(pokemons)      //map[Bulbasaur:Grass/Poison Charmander:Fire Jigglypuff:Normal/Fairy Pikachu:Electric Squirtle:Water]
}
