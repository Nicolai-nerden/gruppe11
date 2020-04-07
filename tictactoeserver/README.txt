
HVORDAN SPILLE PÅ SERVEREN!
For Mac
1. Skriv dette inn I terminalen din og spill i vei:

   nc 178.128.250.190 8080


For windows:
Serveren kommuniserer gjennom netcat slik at man ikke skal trenge å laste ned en klient.
Windows har derimot ikke netcat som en standard. Dette må derfor installeres. Enten på 
homwbrew eller gjennom nmap. Sistnevnte har vi god erfaring med. 

   Nmap kan lastes ned her: https://nmap.org/download.html
   Mer om map: https://en.wikipedia.org/wiki/Nmap

1. Installer netcat, enten det er i homebrew, nmap eller andre program.
2. For nmap kan du skrive dette. 

   Ncat 178.128.250.190 8080


Hvordan man spiller turnering:
Når du starter spiller er det to modus. 
1. PvP - Her skriver man in navnet på alle spillerne og spiller ut turneringen manuelt.
2. Simulasjon - simulering av kvalifisering og enkelt matcher. Du velger selv hvor mange 
   spillere som skal simuleres. Her trenger du kun å trykke på 
   enter for å simulere hele turneringen steg for steg. (Spam enter for å bli kjapt ferdig.)
Hvis 


OM mp01
main.go filen er koden som kjøres på serveren vår. Den bruker pakkene: 
 * client communication (mp01)
 * tictactoeturnering 	(mp01)
 * tictactoe 		(mp01)
 * supportedLangs 	(mp02)
 * validInputd 		(mp02)
Mer informasjon om pakkene finner du på README.txt filen i mp01 folderet.

OM mp02
I makeLangMap er det en applikasjon som kan brukes til å legge til flere skriftspråk på 
tictactoe serveren. 

Her finner du pakkene supportedLangs og validInputs som brukes for å oversette skrifttegn
og andre alternativer for 1-9 om til bytes for å deretter å sjekke om det er et gyldig trekk.   

