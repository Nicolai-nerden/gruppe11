
Hvordan bruke:
1. Åpne terminal i folderet "main".
2. Skriv inn "go run main.go" etterfulgt av to tall som skilles med et mellomrom.

Resultat:
Applikasjonen vil finne ut hva slags type tall som er tastet inn enten det er desimaltall eller hele tall.
Den vil deretter summere de sammen og presentere resultat etterfulgt av hvilken variabeltype det er.

Eksempel på bruk i terminal:
"go run main_sum.go 2.5 3.5"
Terminalen vil gi følgende svar:
"5.7"
"Variable type: float64"

For å analysere med pproff:
1. Åpne main_sum.go i en teksteditor.
2. Fjern // foran "defer profile(...)" øverst i main funksjonen
3. Fjern // foran "github.com/pkg/profile" i import statementet  
4. Skriv inn i terminalen "go build main_sum.go", press enter og deretter./main_sum.go
5. Skriv inn i terminalen "go tool pprof -http=:8080 cpu.pprof"
Dette vil åpne en fane i nettleseren med analyse av filen.

Pass på at du har installert graphviz og Dave Cheney sin profileringspakke.
For å installere skriv in disse linjene i terminalen.
1. "go get github.com/pkg/profile"
2. "brew install graphviz"

(hvis det oppstår feil ved installering av graphviz, pass på at du har installert homebrew. Installasjonguide finner du på nettet.)
