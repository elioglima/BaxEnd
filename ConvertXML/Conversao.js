const fs = require('fs');
const parser1 = new DOMParser();



var parser = require('xml2json');

fs.readFile( './SoapTest.xml', function(err, data) {
    const srcDOM = parser1.parseFromString(data, "application/xml");
    console.log('teste do puag',srcDOM);    
 });



// const retype = require('retypejs');


// let xml = fs.readFileSync('./SoapTest.xml', 'utf-8');
// var convert = require('xml-js');
// var result = convert.xml2json(xml, 
//     {
//         compact: true, 
//         spaces: 4,
//         sanitize:true,
//     });

//     //console.log(result);

//     const model = {
//         "_attributes": "Remove",
//         "Death": "Date",
//         "Inutil": "Remove",
//         "Begin": "DateTime",
//         "End": "DateTime",
//         "Value": "Decimal",
//         "Quantity": "Int"
//     };
    
    
//     var json = JSON.parse(result)

//     let resultClean = retype(json,model);
//     console.log(JSON.stringify(resultClean));





// value = json.GetDadosDevedorResponse.GetDadosDevedorResult.Resultado.Codigo
// console.log(value)

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