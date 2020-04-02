package style

// Style er stilen til htmlfilen
var Style string = `<style>
      html, body {
        width: 100%;
        min-height: 100%;
        overflow-x: hidden;
        margin: 0;
        -webkit-transition: all 0.2s ease 0s;
        transition: all 0.2s ease 0s;
    }
        body {
        font-family: 'roboto', sans-serif;
        font-size: 18px;
        font-weight:900;
        -webkit-font-smoothing: antialiased;
        color: #2c3e50;
        text-align: center;
    }

    main {
        position: relative;
        background-color: lavender;
        margin-right: auto;
        margin-left: auto;
        margin-top: 150px;
        margin-bottom: 50px;
        border-style: solid;
        border-color: white;
        border-radius: 5px;
        color: #333333;
        width: 800px;
        min-height: 300px;
        padding: 50px;
        -webkit-box-shadow: 0px 1px 15px -4px rgba(0,0,0,0.75);
        -moz-box-shadow: 0px 1px 15px -4px rgba(0,0,0,0.75);
        box-shadow: 0px 1px 15px -4px rgba(0,0,0,0.75);
    }

    h1 {
        margin-top: 20px;
    }

    table {
        border-collapse: collapse;
        width:200px;
    }

    th, td {
        text-align: center;
        padding: 8px;
        min-width: 70px;
        font-weight: 500;
    }

    td {
        font-weight: 400;
    }

    tr:nth-child(even){
        background-color: #f2f2f2
    }
    p {
        margin: 15px 3px 12px 3px
    }

    #box {
          position: relative;
          width: 800px;
          min-height: 300px;
          background-color:white;
          display: flex;
          flex-direction: row;
          flex-wrap: wrap;
          justify-content: center;
          align-items: center;
        }
    #kategori {
        height: 100%;
        width: 200px;
        font-weight: 500;
        margin: 5px;
        margin-top: 111px;
        margin-left: 40px;
        margin-bottom: 80px;
        text-align: right;
    }
    #tabell {
        height: 100%;
        width: 400px;
        margin: 5px;
		margin-top: 8px;
    }

  </style>`

// MainStart er starten på tagsene
var MainStart string = `<main>
        <div id="box">
            <h1 style="width: 100%; margin-top: 50px;">Covid-19 i dag:</h1>
            <div id="kategori">
                <p>Smittede totalt:</p>
                <p>Nye smittede idag:</p>
                <p>Døde totalt:</p>
                <p>Nye Døde:</p>
            </div>
            <div id="tabell">
                <table>
                    <tr>
                        <th>Idag</th>
                        <th>Igår</th>
                        <th>Differanse</th>
					</tr>`

// MainEnd er slutten av tagsene
var MainEnd string = `</table>
            </div>
        </div>
    </main>`
