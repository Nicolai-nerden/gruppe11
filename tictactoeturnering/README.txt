
FØR BRUK:
For å bruke denne applikasjonen må du først kopiere folderet "tictactoe" i denne mappen paste det inn i folderet "src" til go på maskinen din.

HVORDAN BRUKE APPLIKASJONEN tictactoeturnering.go:
1. Åpne terminal i denne mappen
2. Skriv "go run tictactoeturnering.go" i terminalen. (Selvom det bare skal spilles en match).
3. Applikasjonen forteller deg hvordan du bruker den.

OM APPLIKASJONEN "tictactoeturnering.go"
Det er to moduser i tictactoeturering.go, PvP og simulasjon.
1. PvP lar deg legge til spillere og starter turneringen deretter. Her blir alle matchene spilt manuelt.
2. Simulasjon simulerer spillere, kvalifiseringsresultater og deretter turneringen. Du må trykke enter-knappen for hvert steg. 
   Alle matcher blir simulert. Hvis du hvil spille ut turneringen kjapt kan du spamme enter-knappen helt til applikasjonen er avsluttet.

OM APPLIKASJONEN "tictactoe.go" i "tictactoe-folderet":
Dette er logikken for selve Tic-Tac-Toe-spillet.
Dette er kun en pakke som brukes i tictactoeturnering.go.
Dette er for å modulere koden i to moduler - en for spillet og en for turneringslogikken.
Hvis du vil spille kun en match må dette likevel gjøres i tictactoeturnering.go

HVORDAN TURNERINGSLOGIKKEN FUNGERER.
1. Spillet starter en kvalifiseringsrunde med alle spillerne lagt til.
2. Alle spiller dermed to matcher hver. 
3. Det er ikke mulig å få uavgjort. Hvis spillet blir fylt uten en vinner, blir vinneren avgjort basert på hvem som har brukt minst tid.
4. Spillet sorterer tabellen etter hvem som har høyest score. Disse er igjen sortert etter hvem som har brukt minst tid totalt i kvalifiseringsrunden. 
   Slik er de med mest poeng øverst og i hver poengkategori er disse sortert etter minst tid brukt.
5. Det øverste antallet spillerne som er en eksponent av 2 (2^n) dvs. 2, 4, 8, 16 osv. Kvalifiserer seg til sluttspillet.
6. Sluttspillet er et tradisjonell turnering der vinneren går videre til neste turneringsrunde. 
7. Til slutt er det en finale i turneringen. Vinneren av denne har vunnet hele turneringen.









