
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
1. Last ned Libterm for iOS eller Terminal Emulator for android.
2. Libterm: "nc 178.128.250.190 8080"
   Terminal Emulator: "toybox nc 178.128.250.190 8080"

LibTerm for Android: https://baixarapk.gratis/en/app/1380911705/libterm
Terminal emulator: https://play.google.com/store/apps/details?id=jackpal.androidterm&hl=en%5D%5B1%5D


For windows:
Serveren kommuniserer gjennom netcat slik at man ikke skal trenge å laste ned en klient.
Windows har derimot ikke netcat som en standardkommando. Vi har derfor laget vår egen 
Klient for windowsbrukere som du finner i clientfolderet. Den er hardcodet til å kobles til den 
eksterne serveren vår. 
Hvis du vil endre ip og port til å kjøre som client til main.go lokalt.
Bytt ut linje 12 i koden med dette: 

conn, err := net.Dial("tcp4", "127.0.0.1:8081")



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

