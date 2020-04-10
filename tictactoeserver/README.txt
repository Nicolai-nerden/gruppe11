
HVORDAN SPILLE PÅ SERVEREN!
For Mac
1. Skriv dette inn I terminalen din og spill i vei:

   nc 178.128.250.190 8080


For windows:
Serveren kommuniserer gjennom netcat slik at man ikke skal trenge å laste ned en klient.
Windows har derimot ikke netcat som en standard. Dette må derfor installeres. Enten på 
homwbrew eller gjennom nmap. Sistnevnte har vi god erfaring med. 

   Nmap kan lastes ned her: https://nmap.org/download.html
   Mer om Nmap: https://en.wikipedia.org/wiki/Nmap

1. Installer netcat, enten det er i homebrew, nmap eller andre program.
2. For nmap kan du skrive dette. 

   Ncat 178.128.250.190 8080

HVIS SERVEREN ER NEDE:
Hvis nc 178.128.250.190 8080 ikke fungerer, betyr dette at serveren har krasjet. Si gjerne i fra
til en av medlemmene i gruppe11. 
Du kan likevel bruke applikasjonen lokalt.

1. Åpne terminal i dette folderet.
2. kjør "go run main.go"
3. Åpne ny terminal
4. kjør "nc localhost 8081" i den nye terminalen.
   (Ncat localhost 8080 for windows)

OM tictactoeserver
main.go filen er koden som kjøres på serveren vår. Den bruker diverse pakker vi har lagd.
* I hver i turnering, tictactoe, og comm er det to nesten identiske pakker.
* Grunnen til at det er to nesten like pakker er fordi serveren bruker to ulike måter å
  kommunisere avhengig av om den spiller lokalt eller multiplayer.
* Det er forskjell på funksjonalitet i lokal og multiplayer for turneringlogikken. I 
  lokal. I Lokal kan du simulere spill. Det kan du ikke i multiplayer. Det er andre små
  Forskjeller også.


Mer informasjon om pakkene finner du på ABOUT.txt.

OM mp02
I makeLangMap er det en applikasjon som kan brukes til å legge til flere skriftspråk på 
tictactoe serveren. 

Her finner du pakkene supportedLangs og validInputs som brukes for å oversette skrifttegn
og andre alternativer for 1-9 om til bytes for å deretter å sjekke om det er et gyldig trekk.   

