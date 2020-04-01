package supportedLangs

// ValidInputs lagrer de gyldige inputene til hvert språk.
// Funksjonelt er det en slice som inneholder maps med int som key og slice av bytes som value.
var ValidInputs = []map[int][]byte{ // legg til navnet på språket laget på en egen linje, husk komma på slutten. Kommenter navnet ditt bak.
	standard,      // Nicolai
	mandarin,      // Nicolai
	binary,        // Nicolai
	romanNumerals, //Benjamin
	korean, //Mati
}

// SupportedLangs inneholder navnet på de ulike språkene som er støttet.
var SupportedLangs = []string{ // Skriv in navnet på skriftspråket på norsk. Husk komma etter!
	"Vanlige tall",
	"Mandarin",
	"Binære tall (4 bits)",
	"Romertall",
	"Koreansk"
}

// lim in maps med støddede språk under hverandre her
var standard = map[int][]byte{
	1: {49},
	2: {50},
	3: {51},
	4: {52},
	5: {53},
	6: {54},
	7: {55},
	8: {56},
	9: {57},
}

var mandarin = map[int][]byte{
	1: {228, 184, 128},
	2: {228, 186, 140},
	3: {228, 184, 137},
	4: {229, 155, 155},
	5: {228, 186, 148},
	6: {229, 133, 173},
	7: {228, 184, 131},
	8: {229, 133, 171},
	9: {228, 185, 157},
}

var binary = map[int][]byte{
	1: {48, 48, 48, 49},
	2: {48, 48, 49, 48},
	3: {48, 48, 49, 49},
	4: {48, 49, 48, 48},
	5: {48, 49, 48, 49},
	6: {48, 49, 49, 48},
	7: {48, 49, 49, 49},
	8: {49, 48, 48, 48},
	9: {49, 48, 48, 49},
}

var romanNumerals = map[int][]byte{
	1: {73},
	2: {73, 73},
	3: {73, 73, 73},
	4: {73, 86},
	5: {86},
	6: {86, 73},
	7: {86, 73, 73},
	8: {86, 73, 73, 73},
	9: {73, 88},
}

var korean = map[int][]byte{
	1: {237, 149, 152, 235, 130, 152},
	2: {235, 145, 152},
	3: {236, 133, 139},
	4: {22, 235, 132, 183},
	5: {235, 139, 164, 236, 132, 175},
	6: {236, 151, 172, 236, 132, 175},
	7: {236, 157, 188, 234, 179, 177},
	8: {236, 151, 172, 235, 141, 159},
	9: {236, 149, 132, 237, 153, 137},
}
