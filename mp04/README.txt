
SERVEREN FOR mp04:

http://165.22.198.8:8080/

OM PAKKEN main
1. Applikasjonen finner de to siste rapportene som er lagt ut a who. (Package latestreportlinks)
2. Den analyserer dermed disse PDF-filene for å finne norges statistikk. (package covidanalytics)
3. Den oppretter dermed en local-server og viser frem statistikken sammenlignet i nettleser. (Package style)

OM PAKKEN covidanalytics
*Dette er en package som brukes i mp04.go
Applikasjonen analyserer og sammenligner to covid-rapporter og returnerer statistikken.

OM PAKKEN latestreportlinks
*Dette er en package som brukes i covidanalytics.go
Applikasjonen søker gjennom nettsiden:
https://www.who.int/emergencies/diseases/novel-coronavirus-2019/situation-reports
Og finner de to siste rapportene. (Kan programmeres om til å finne to spesifikke rapporter.)
Den returnerer de to linkene til filene tilbake til covidanalytics.go

OM PAKKEN style
* Dette er en package som brukes i mp04.go
Den inneholder kun strings med html og css kode for styling av nettsiden.

