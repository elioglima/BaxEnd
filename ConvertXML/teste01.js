const parseString = require('xml2js').parseString;
const fs = require('fs')
const retype = require('retypejs');
const xml = fs.readFileSync('./SoapTest.xml');



    const model = {
        "$": "Remove",
        "Death": "Date",
        "Inutil": "Remove",
        "Begin": "DateTime",
        "End": "DateTime",
        "Value": "Decimal",
        "Quantity": "Int"
    };



parseString(xml, function (err, result) {
    const a = JSON.stringify(result);
    const b = JSON.parse(a);
    


console.log(result.GetDadosDevedorResponse.GetDadosDevedorResult[0].Resultado)

    console.dir(JSON.stringify(retype(b, model)));
});