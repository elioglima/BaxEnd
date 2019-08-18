const soapRequest = require('./SoapRequest');
const fs = require('fs');
const xml2js = require('xml2js')
const url_homologacao = "https://acordocertohomo.mfmti.com.br/WebService.asmx?WSDL"
const url_producao    = "https://integracaonegociafacil.mfmti.com.br/?wsdl"

GetValue = (objeto, tagName) => {
  if (objeto == null) {
    return null
  } else if (objeto[tagName] === "undefined") {
    return null
  }
  
  objeto = objeto[tagName]        
  if (Array.isArray(objeto)) {
    if (objeto.length === 0) {            
      objeto = null
    } else if (objeto.length === 1) {
      objeto = objeto[0]            
    } 
  }
  return objeto
}

class GetDadosDevedorResult {
  constructor() {
    this.Codigo = ''
    this.Data = ''
  }

  Parse(xml) {
    let xml2js = require('xml2js');
    let parser = new xml2js.Parser();

    return new Promise(function(resolve, reject) {
      parser.parseString(xml, function(err,result){                        
        
        if (result === null) {
          reject({Status:500,Response:"Servidor não respondeu."})
        } 

        let objeto = GetValue(result, "soap:Envelope")
        if (objeto === null) {
          reject({Status:500,Response:"Tag soap:Envelope não localizada."})
          return 
        }

        objeto = GetValue(objeto, "soap:Body")
        if (objeto === null) {
          reject({Status:500,Response:"Tag 'soap:Body' não localizada."})
          return 
        }

        objeto = GetValue(objeto, "GetDadosDevedorResponse")
        if (objeto === null) {
          reject({Status:500,Response:"Tag 'GetDadosDevedorResponse' não localizada."})
          return 
        }

        objeto = GetValue(objeto, "GetDadosDevedorResult")
        if (objeto === null) {
          reject({Status:500,Response:"Tag 'GetDadosDevedorResult' não localizada."})
          return 
        }

        objeto = GetValue(objeto, "Resultado")
        if (objeto === null) {
          reject({Status:500,Response:"Tag 'Resultado' não localizada."})
          return 
        }

        objeto = GetValue(objeto, "Codigo")
        if (objeto === null) {
          reject({Status:500,Response:"Tag 'GetDadosDevedorResult' não localizada."})
          return 
        }

        resolve({Status:200,Response:"Dados carregados com sucesso.",Objeto:objeto})
      });
    });
  }
}


execute = () => {  
    
    // console.log('Inicio Execute')    
    const url = url_producao;
    const headers = {
      'user-agent': 'sampleTest',
      'Content-Type': 'text/xml;charset=UTF-8',
    //   'soapAction': 'https://graphical.weather.gov/xml/DWMLgen/wsdl/ndfdXML.wsdl#LatLonListZipCode',
    };
    

    const xml = fs.readFileSync('./XmlEnvio.xml', 'utf-8');
    (async () => {
      const { response } = await soapRequest(url, headers, xml, 1000)
      const { body, statusCode } = response
      // console.log("statusCode",statusCode)    
      // console.log("body",body)         

      GetDadosDevedorResult = new GetDadosDevedorResult()
      GetDadosDevedorResult.Parse(body)
      .then(
        dados => {
          console.log("then",dados)
        }
      ).catch(
        dados => {
          console.log("catch",dados)
        }
      );

    })();
};

execute()