
VIDEOEKSEMPEL AV APPLIKASJONEN I BRUK:
I dette folderet er det en lenkefil. Denne åpner en video som viser et praktisk eksempel
av spill på serveren fra en iPhone via terminal-applikasjonen LibTerm
Lenkefilen fører til: 
https://drive.google.com/file/d/1mEcprU-Io0QhYD2c_NAMLEUFgzcjMJnR/view

FOR SENSOR / JANIS / HJELPELÆRERE
Hvis du ikke får spilt multiplayer kan du skrive "forcerestart" f.eks når den ber deg om navn.
(Hvis du vurderer oppgaven vår er det bare å gjøre dette selvom en turnering evt. er i gang.) 

FOR ALLE ANDRE
Hvis du kobles på serveren, og du ikke får spilt multiplayer OG du vet dette ikke er fordi
noen andre spiller kan du skrive "forcerestart" for å restarte serveren. 

HVORDAN SPILLE PÅ SERVEREN!
For Mac
1. Skriv dette inn I terminalen din og spill i vei:

   nc 178.128.250.190 8080

For Mobile enheter
1. Last ned Libterm for iOS-eller Android.
2. Se steg 1 "For Mac"

LibTerm for Android: https://baixarapk.gratis/en/app/1380911705/libterm
LibTerm for iOS: https://apps.apple.com/us/app/libterm/id1380911705?ign-mpt=uo%3D4


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
Du kan likevel bruke serveren lokalt.

1. Åpne terminal i dette folderet.
2. kjør "go run main.go"
3. Åpne ny terminal
4. kjør "nc localhost 8081" i den nye terminalen.
   (Ncat localhost 8080 for windows)

OM tictactoeserver
main.go filen er koden som kjøres på serveren vår. Den bruker diverse pakker vi har lagd.
* I turnering, "tictactoe", og comm er det lignende pakker en for multiplayer og en for lokalt spill via serveren.
* Grunnen til dette er at serveren bruker to ulike måter å
  kommunisere avhengig av om den spiller lokalt eller multiplayer.
* Det er forskjell på funksjonalitet i lokal og multiplayer for turneringlogikken. I Lokalkan du simulere spill. 
  Det kan du ikke i multiplayer. Det er andre små forskjeller også slik som å legge til spillere manuelt og velge 
  Mellom simulasjonsmodus og pvp. 

I pakken makeLangMap i utf er det en applikasjon som kan brukes til å legge til flere skriftspråk på 
tictactoe serveren. Denne er relatert til mp02. 

I pakken utf finner du også pakkene supportedLangs og validInputs som brukes for å oversette skrifttegn
og andre alternativer for 1-9 om til bytes for å deretter å sjekke om det er et gyldig trekk.   

Mer informasjon om pakkene finner du på ABOUT.txt.

