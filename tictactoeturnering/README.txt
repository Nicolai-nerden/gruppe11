
HVORDAN KJØRE APPLIKASJONEN:
macOS/Linux: kjør "tictactoeturnering"

Hvis den ikke vil kjøre:
1. Åpne terminal og skriv "chmod +x " (ikke trykk enter enda)
2. Dra inn tictactoeturnering-filen i terminalen og slipp. (Ikke .go eller .exe filen)
3. Trykk enter. Filen vil nå tolkes som en exec fil og ikke et dokument.
4. Kjør filen. 

Hvis den ikke vil kjøre fordi det er fra en uidentifisert utvikler:
1. Åpne "system preferences" 
2. Gå til "Security & Privacy"
3. Trykk knappen "Open anyway"

Windows: kjør "tictactoeturnering.exe"

Applikasjonen forteller deg hvordan du bruker den.


HVIS DU VIL GJØRE ENDRINGER I tictactoeturnering.go OG KJØRE DEN MED NYE ENDRINGER:
1. Kopier folderet "tictactoe" i denne mappen lim det inn i folderet "src" til go på maskinen din.
2. Gjør endringer i valgfri tekstbehandler f.eks. VS Code.
3. Åpne terminal i denne mappen.
4. Skriv "go run tictactoeturnering.go" i terminalen.


OM APPLIKASJONEN "tictactoeturnering.go"
Det er to moduser i tictactoeturering.go, PvP og Simulasjon.
1. PvP lar deg legge til spillere og starter turneringen deretter. Her blir alle matchene spilt manuelt.
2. Simulasjon simulerer spillere, kvalifiseringsresultater og deretter turneringen. Du må trykke enter-knappen for hvert steg. 
   Alle matcher blir simulert. Hvis du hvil spille ut turneringen kjapt kan du spamme enter-knappen helt til applikasjonen er avsluttet.


OM APPLIKASJONEN "tictactoe.go" i "tictactoe-folderet":
Dette er logikken for selve Tic-Tac-Toe-spillet.
Dette er kun en pakke som brukes i tictactoeturnering.go.
Dette er for å modulere koden i to moduler - en for spillet og en for turneringslogikken.
Hvis du vil spille kun en match må dette likevel gjøres i tictactoeturnering-applikasjonen.


HVORDAN TURNERINGSLOGIKKEN FUNGERER.
1. Hvis antallet spillere er en toerpotens (2^n - dvs. 2, 4, 8, 16 osv.) Starter den sluttspill med engang.
   Hvis ikke starter spillet starter kvalifiseringsmatcher med alle spillerne lagt til 
2. I kvalifiseringen spiller alle spillerne to matcher hver. En som X og en som O
3. Det er ikke mulig å få uavgjort. Hvis spillet blir fylt uten en vinner, blir vinneren avgjort basert på hvem som har brukt minst tid.
4. Etter kvalifiseringen sorterer tabellen etter hvem som har høyest score. Disse er igjen sortert etter hvem som har brukt minst tid totalt i kvalifiseringsrunden. 
   Slik er de med mest poeng øverst og i hver poengkategori er disse sortert etter minst tid brukt.
5. Det øverste antallet spillerne som er en toerpotens (2^n). Kvalifiserer seg til sluttspillet.
6. Sluttspillet er et tradisjonell turnering der vinneren går videre til neste turneringsrunde. 
7. Til slutt er det en finale i turneringen. Vinneren av denne har vunnet hele turneringen.









