
OM MAPPENE:

Dette er kun pkg som brukes i "tictactoe.go" og "tictactoeturnering.go" i mp01 mappen.

HVORDAN LEGGE TIL FLERE SKRIFTSPRÅK:
Videoen som er linket i denne mappen viser hvordan.
(https://drive.google.com/file/d/1noOXmF83O5XRLAz_KioY2Vu_eeeqBZVs/view?usp=sharing)

1. Bruk applikasjonen "makeLangMap.go"
2. Kopier utskriften når applikasjonen er ferdig.
3. Lim det inn i supportedLangs.go og legg til norsk navn og navnet på slicen i ValidInputs sliten i denne filen.

OM makeLangMap
Denne konverterer tegnene du taster inn til bytes og printer ut ett stykke kode som du kan lime inn i supportedLangs.go. 
Slik kan du legge til nye skriftspråk.

OM SupportedLangs:
*Dette er en package som brukes av validInputs.go
Inneholder tabeller med ulike språk og deres skrifttegn 1-9 i bytes (values) og deres tilhørende verdi i form av vanlige tall (keys).


Om validInputs:
*Dette er en package som brukes av tictactoeturnering.go og tictactoe.go
Den sjekker om inputs av ulike skrifttegn matcher med 1-9 i de språkene vi har lagt til.
Returnerer det samhørende tallet hvis det er støttet. Ber deg om å prøve igjen om har tastete ugyldig input eller et tegn som ikke er støttet.