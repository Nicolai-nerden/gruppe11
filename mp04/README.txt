
OM mp04.go
1. Applikasjonen finner de to siste rapportene som er lagt ut a who.
2. Den analyserer dermed disse PDF-filene for å finne norges statistikk.
3. Den oppretter dermed en local-server og viser frem statistikken sammenlignet i nettleser.

OM covidanalytics.go
*Dette er en package som brukes i mp04.go
Applikasjonen analyserer og sammenligner to covid-rapporter. Og returnerer statistikken.

OM latestreportlinks.go
*Dette er en package som brukes i covidanalytics.go
Applikasjonen søker gjennom nettsiden:
https://www.who.int/emergencies/diseases/novel-coronavirus-2019/situation-reports
Og finner de to siste rapportene. (Kan programmeres om til å finne to spesifikke rapporter.)
Den returnerer de to linkene til filene til covidanalytics.go

OM style.go
* Dette er en package som brukes i mp04.go
Den inneholder kun strings med html og css kode for styling av nettsiden.

