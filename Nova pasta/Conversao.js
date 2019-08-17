const fs = require('fs');

let xml = fs.readFileSync('./SoapTest.xml', 'utf-8');
var convert = require('xml-js');
var result = convert.xml2json(xml, 
    {
        compact: true, 
        spaces: 4,
        sanitize:true,
    });
var json = JSON.parse(result)

value = json.GetDadosDevedorResponse.GetDadosDevedorResult.Resultado.Codigo
console.log(value)

/*

<GetDadosDevedorResponse xmlns="http://tempuri.org/">
    <GetDadosDevedorResult>
        <Resultado xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns="">
            <Codigo>00</Codigo>
            <Data>16/08/2019 18:45:31</Data>
        </Resultado>
    </GetDadosDevedorResult>
</GetDadosDevedorResponse>


*/