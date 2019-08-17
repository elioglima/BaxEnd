const soapRequest = require('../../libs/SoapRequest');
const fs = require('fs');

const url_homologacao = "https://acordocertohomo.mfmti.com.br/WebService.asmx?WSDL"
const url_producao    = "https://integracaonegociafacil.mfmti.com.br/?wsdl"


module.exports.execute = (req, res) => {  
    
    console.log('Inicio Execute')    
    const url = url_producao;
    const headers = {
      'user-agent': 'sampleTest',
      'Content-Type': 'text/xml;charset=UTF-8',
    //   'soapAction': 'https://graphical.weather.gov/xml/DWMLgen/wsdl/ndfdXML.wsdl#LatLonListZipCode',
    };
    

    const xml = fs.readFileSync('c:/Teste/SoapTest.xml', 'utf-8');
    (async () => {

      const { response } = await soapRequest(url, headers, xml, 1000)
      const { body, statusCode } = response
      console.log("statusCode",statusCode)    

      var convert = require('xml-js');

    //   <?xml version="1.0" encoding="utf-8"?>
    // <GetDadosDevedorResponse xmlns="http://tempuri.org/"><GetDadosDevedorResult><Resultado xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns=""><Codigo>00</Codigo><Data>16/08/2019 18:45:31</Data></Resultado></GetDadosDevedorResult></GetDadosDevedorResponse></soap:Body></soap:Envelope>
    
    // var xmls = body.replace('<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">','')
    // xmls = xmls.replace('<?xml version="1.0" encoding="utf-8"?>','')
    // xmls = xmls.replace('<soap:Body>','')
    // xmls = xmls.replace(' xmlns="http://tempuri.org/"','')
    // xmls = xmls.replace(' xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns=""','')
    // xmls = xmls.replace('</soap:Body>','')
    // xmls = xmls.replace('</soap:Envelope>','')

    var fs = require('fs');
var parse = require('xml-parser');
var inspect = require('util').inspect;
 
var obj = parse(body);
json = inspect(obj, { colors: true, depth: Infinity })
console.log(json);

      
    // console.log('')
    // console.log('xmls', xmls)
    // console.log('\n', result)
    


    })();
};