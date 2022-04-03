package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var dob int
	var name string
	var surname string

	dobFlag := flag.Int("day", 0, "Day of birth")
	nameFlag := flag.String("name", "unset", "Name")
	surnameFlag := flag.String("surname", "unset", "Surname")
	flag.Parse()

	if *dobFlag == 0 {
		fmt.Println("Day not detected\nEnter day of your birth")
		fmt.Scanf("%d", &dob)
	} else {
		dob = *dobFlag
	}

	if *nameFlag == "unset" {
		fmt.Println("Name not detected\nEnter your name")
		fmt.Scanf("%s\n", &name)
	} else {
		name = *nameFlag
	}

	if *surnameFlag == "unset" {
		fmt.Println("Surname not detected\nEnter your surname")
		fmt.Scanf("%s\n", &surname)
	} else {
		surname = *surnameFlag
	}

	nickname := dateOfBirthMap[dob] +
		nameMap[strings.ToLower(string(name[0:1][0]))] +
		surnameMap[strings.ToLower(string(surname[0:1][0]))]
	fmt.Println(nickname)
}

var dateOfBirthMap map[int]string = map[int]string{
	1: "Król", 2: "Senpai", 3: "Mistrz", 4: "Drwal", 5: "Ksiądz", 6: "Prezes", 7: "Chojrak",
	8: "Lord", 9: "Ojciec", 10: "Kaczka",
	11: "Mnich", 12: "Wieloryb", 13: "Wielbłąd", 14: "Samuraj", 15: "Towarzysz", 16: "Wódz", 17: "Alkoholik",
	18: "Knur", 19: "Locha", 20: "Kot",
	21: "Raper", 22: "Królewna", 23: "Rolnik", 24: "Kolega", 25: "Sir", 26: "Brzydal", 27: "Przystojniak",
	28: "Tyran", 29: "Anżej", 30: "Oszołom", 31: "Zombie",
}
var nameMap = map[string]string{
	"a": " W Niedoli ", "b": " W Biedzie ", "c": " W Wannie ", "d": " W Kościele ", "e": " W Azji ",
	"f": " We Francji ", "g": " W konsternacji ", "h": " Na Hamaku ", "i": " W Iglo ", "j": " Wielki ",
	"k": " Zacny ", "l": " Leniwy ", "m": " Mądry ", "n": "Niedorozwinięty", "o": "Ordynarny", "p": "Prawy",
	"r": " Ryczący ", "s": "Sprytny", "t": " W Tempurze ", "u": " W Gościnie ", "w": " W Wannie", "z": " W Zagrodzie ",
}
var surnameMap map[string]string = map[string]string{
	"a": " Z Kaczką ", "b": " Z Siekierą W Głowie ", "c": " Z Zamiarem Zabicia ", "d": " Z dziwną miną ",
	"e": " O brzydkiej twarzy ", "f": " Z Wieloma Pryszczami ", "g": " Z Garbem ", "h": " Pełen Nadzei ",
	"i": " Bez Lewej stopy ", "j": " Z Hymnem na ustach ", "k": " Z Zezem Rozbieżnym ", "l": " Rozczłonkowany ",
	"m": " O Karlim Wzroście ", "n": " O Nielicznych talentach ", "o": " Z Koleżanką Kolegi ", "p": " Prawie żywy",
	"r": " Z Rymowanym Powiedzonkiem ", "s": " O szerokim czole ", "t": " Z Twardą Makówką ",
	"u": " Z kieliszkiem przy ustach ", "w": " Z Małą Wiedzą ", "z": " Z Zamkniętymi Powiekami ",
}
