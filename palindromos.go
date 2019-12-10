package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Este programa requiere que el archivo de entrada tenga los strings separados por salto de linea.
//En caso de tenerotra regla para separar, hay que modificar la linea 66

// 2. De una lista de input de strings que va a venir en un archivo llamado "palabrotas.txt",
// indicar cuales de ellos son palíndromos y cuales no. Indicar cuantas veces se repite cada palabra.
// [Variación]: Ignorar todo lo que no sea un caracter alfanumérico.
func rstring(s string) string {
	runes := []rune(s)
	n := len(runes) - 1

	for i := 0; i < n; i++ {
		runes[i], runes[n] = runes[n], runes[i]
		n = n - 1
	}

	return string(runes)
}

func ignore(s string) string {
	runes := []rune(s)
	var ignored []rune

	for i := 0; i < len(runes); i++ {
		if runes[i] >= 48 && runes[i] <= 57 { //NUMBERS
			ignored = append(ignored, runes[i])
		}
		if runes[i] >= 97 && runes[i] <= 122 { //LowerCase
			ignored = append(ignored, runes[i])
		}
		if runes[i] >= 65 && runes[i] <= 90 { //UpperCase
			ignored = append(ignored, runes[i])
		}
		if runes[i] >= 193 && runes[i] <= 221 { //accented LowerCase
			ignored = append(ignored, runes[i])
		}
		if runes[i] >= 225 && runes[i] <= 253 { //accented UperCase
			ignored = append(ignored, runes[i])
		}
	}

	// for i := 0; i < len(runes); i++ {
	// 	if unicode.IsLetter(runes[i]) || unicode.IsNumber(runes[i]) { //NUMEROS
	// 		ignored = append(ignored, runes[i])
	// 	}
	// }

	return string(ignored)
}

func main() {
	fmt.Println("Hola Mundo 🌍!")
	//Empiezo yo
	path := "palabrotas.txt"
	byteArray, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Archivo inválido en la ruta: '", path, "' ", err.Error())
		os.Exit(1)
	}
	content := string(byteArray)

	stringArray := strings.Split(content, "\n") //Elimino los saltos de linea

	myMap := make(map[string]int)

	for i := 0; i < len(stringArray); i++ {
		key := stringArray[i]
		key = ignore(key)                        //Calls the ignore function
		if value, exists := myMap[key]; exists { //If repeated add 1
			myMap[key] = value + 1
		} else {
			myMap[key] = 1 //If not, initialize
		}

	}
	for key, value := range myMap { //Print the results
		fmt.Println("🐀🐀🐀🐀🐀🐀🐀🐀🐀🐀🐀🐀🐀🐀🐀🐀🐀🐀🐀🐀🐀🐀🐀")
		if value == 1 {
			fmt.Println(key, "no se repite en el archivo")
		} else {
			fmt.Println(key, " se repite en el archivo ", value-1, " veces")
		}
		ikey := rstring(key)       //rstring verifica si es palindromo o no
		key = strings.ToLower(key) //Para que no distinga mayusculas y minusculas
		ikey = strings.ToLower(ikey)
		if key == ikey {
			fmt.Println("Ademas, es palíndromo.")
		}
	}

	//Experimento exitoso con rune
	modif := "0"
	runes := []rune(modif)
	fmt.Println(modif, runes)
	modif = "9"
	runes = []rune(modif)
	fmt.Println(modif, runes)
	modif = "a"
	runes = []rune(modif)
	fmt.Println(modif, runes)
	modif = "z"
	runes = []rune(modif)
	fmt.Println(modif, runes)
	modif = "A"
	runes = []rune(modif)
	fmt.Println(modif, runes)
	modif = "Z"
	runes = []rune(modif)
	fmt.Println(modif, runes)
	modif = "Á"
	runes = []rune(modif)
	fmt.Println(modif, runes)
	modif = "Ý"
	runes = []rune(modif)
	fmt.Println(modif, runes)
	modif = "á"
	runes = []rune(modif)
	fmt.Println(modif, runes)
	modif = "ý"
	runes = []rune(modif)
	fmt.Println(modif, runes)
	modif = "Ñ"
	runes = []rune(modif)
	fmt.Println(modif, runes)
	modif = "ñ"
	runes = []rune(modif)
	fmt.Println(modif, runes)
}
