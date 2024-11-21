package main

import (
	"bufio"
	"fmt"
	"maps"
	"math/rand"
	"os"
	"strings"
	"time"
)

var counting = map[int]string{
	1:  "eradi",
	2:  "ischinaim",
	3:  "islocha",
	4:  "arbaa",
	5:  "ramichaa",
	6:  "chichaa",
	7:  "chivaa",
	8:  "chimonaa",
	9:  "tichaa",
	10: "assaraa",
}

var tens = map[int]string{
	20: "esrim",
	30: "islochim",
	40: "arbaim",
	50: "ramichim",
}

var rules = map[int]string{
	1: "menor caminho, máxima velocidade",

	2: "analise sempre o ambiente, use-o a seu favor, visualize rotas de fuga e possíveis atacantes",

	3: "toda defesa exige um ataque simultâneo",

	4: "a surpresa, a dissimulação e a explosão no ataque são as melhoras armas do krav Maga",

	5: "no krav Maga nunca recuamos",

	6: "a mente sempre deve visualizar uma situação real",

	7: "nunca subestime seu inimigo, ele pode saber mais que você",

	8: "ao entrar em um conflito, saia o mais rapidamente possível dele",

	9: "nunca pare até aniquilar toda a ameaça",

	10: "controle sua mente, pois em uma situação de crise, não ganha quem é mais forte e sim quem mais tem sangue frio e sabe controlar suas emoções",
}

var randomizer *rand.Rand

func main() {
	s := rand.NewSource(time.Now().UnixNano())
	randomizer = rand.New(s)
	for {
		text := getInput(`
==============================
= What do you want to test?  =
==============================
1 - Numbers (int to str)
2 - Numbers with tens (int to str)
3 - Numbers (int to str)
4 - Numbers with tens (int to str)
7 - Rules (int to str)
8 - Rules (str to int)
0 - Exit
Your option: `)
		text = strings.Trim(text, "\n")
		switch text {
		case "1":
			testNumbers(false)
		case "2":
			testNumbers(true)
		case "3":
			testNumberNames(false)
		case "4":
			testNumberNames(true)
		case "7":
			testRules()
		case "8":
			testRuleNames()
		default:
			fmt.Println("\n ---= Finished Execution =--- \n------------")
			return
		}
	}
}

func testNumbers(includeTens bool) {
	numbers := map[int]string{}
	maps.Copy(numbers, counting)
	if includeTens {
		maps.Copy(numbers, tens)
	}
	for {
		res := randItem(numbers)
		input := getInput(fmt.Sprintf("\nNumero: %d\nNome?: ", res.int))
		if input == "0" {
			return
		}
		if input == res.string {
			fmt.Println("Correto ✅")
		} else {
			fmt.Println("Errado ❌ -> ", res.string)
		}
	}
}

func testNumberNames(includeTens bool) {
	numbers := map[int]string{}
	maps.Copy(numbers, counting)
	if includeTens {
		maps.Copy(numbers, tens)
	}
	for {
		res := randItem(numbers)
		input := getInput(fmt.Sprintf("\nNome: %s\nNumero?: ", res.string))
		if input == "0" {
			return
		}
		if input == fmt.Sprintf("%d", res.int) {
			fmt.Println("Correto ✅")
		} else {
			fmt.Println("Errado ❌ -> ", res.int)
		}
	}
}

func testRules() {
	for {
		res := randItem(rules)
		input := getInput(fmt.Sprintf("\nNumero: %d\nRegra?: ", res.int))
		if input == "0" {
			return
		}
		str := res.string
		str = strings.ReplaceAll(str, "ê", "e")
		str = strings.ReplaceAll(str, ",", "")
		str = strings.ReplaceAll(str, ".", "")
		str = strings.ReplaceAll(str, "-", "")
		str = strings.ReplaceAll(str, "é", "e")
		str = strings.ReplaceAll(str, "ã", "a")
		str = strings.ReplaceAll(str, "â", "a")
		str = strings.ReplaceAll(str, "í", "i")
		if input == strings.ToLower(str) {
			fmt.Println("Correto ✅")
		} else {
			fmt.Println("Errado ❌ -> ", str)
		}
	}
}

func testRuleNames() {
	for {
		res := randItem(rules)
		input := getInput(fmt.Sprintf("\nRegra: %s\nNumero?: ", res.string))
		if input == "0" {
			return
		}
		if input == fmt.Sprintf("%d", res.int) {
			fmt.Println("Correto ✅")
		} else {
			fmt.Println("Errado ❌ -> ", res.int)
		}
	}
}

func getInput(text string) (result string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(text)
	result, _ = reader.ReadString('\n')
	result = strings.Trim(result, "\n")
	result = strings.ToLower(result)
	result = strings.ReplaceAll(result, "ê", "e")
	result = strings.ReplaceAll(result, ",", "")
	result = strings.ReplaceAll(result, ".", "")
	result = strings.ReplaceAll(result, "-", "")
	result = strings.ReplaceAll(result, "é", "e")
	result = strings.ReplaceAll(result, "ã", "a")
	result = strings.ReplaceAll(result, "â", "a")
	result = strings.ReplaceAll(result, "í", "i")
	// result = strings.ReplaceAll(result, "", "")
	// result = strings.ReplaceAll(result, "", "")

	return
}

func randItem(items map[int]string) struct {
	int
	string
} {
	pos := randomizer.Intn(len(items)) + 1
	if pos == 11 {
		pos = 20
	} else if pos == 12 {
		pos = 30
	} else if pos == 13 {
		pos = 40
	} else if pos == 14 {
		pos = 50
	}
	return struct {
		int
		string
	}{
		pos,
		items[pos],
	}
}
