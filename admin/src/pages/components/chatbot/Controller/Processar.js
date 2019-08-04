
export const Pesquisa = [
    {
        "id":1,
        "Titulo":"Hoje é: {datetime}",
        "Indexs":[
            {"chave":"Que dia é hoje"},
            {"chave":"Qual dia é hoje"},
            {"chave":"Hoje é"},
            {"chave":"Hoje é"},
        ]
    },{
            "id":1,
            "Titulo":"Meu nome é Davi, sou um robô, estou aqui para te ajuda e é um prazer em te conhecer.",
            "Indexs":[
                {"chave":"Qual é seu nome"},
                {"chave":"Qual e seu nome"},
                {"chave":"qual seu nome"},
                {"chave":"Com quem estou falando"},
                {"chave":"Como você se chama"},
                {"chave":"Quem é você"},
            ]
    },{
        "id":1,
        "Titulo":"É um prazer te ajudar, caso precise é só chamar.",
        "Indexs":[
            {"chave":"origado"},
        ]
    },{
        "id":1,
        "Titulo":"Acesso ao Sistema",
        "Indexs":[
            {"chave":"acesso sistema"},
            {"chave":"acesso de sistema"},
            {"chave":"acesso ao sistema"},
            {"chave":"acesso"},
            {"chave":"sistema"}
        ]
    },
    {
        "id":2,
        "Titulo":"Manual do Usuário",
        "Indexs":[
            {"chave":"manual usuario"},
            {"chave":"manual de usuario"},
            {"chave":"manual do usuario"},
            {"chave":"manual"},
            {"chave":"usuario"}
        ]
    },
    {
        "id":3,
        "Titulo":"Licença de Uso",
        "Indexs":[
            {"chave":"uso"},
            {"chave":"licenca uso"},
            {"chave":"licença uso"},
            {"chave":"licença de uso"},
            {"chave":"licenca"},
            {"chave":"licença"},
            {"chave":"autorização"},
            {"chave":"autorizacao"}
        ]
    }

];


export const chatbot = [
{
    "id":1,
    "msg":"Olá, eu sou o Davi e estou aqui para lhe ajudar.",
    "temp":1,        
},
{
    "id":2,
    "msg":"Digite em poucas palavras o que deseja.",
    "temp":1,        
},
];

const procComparativoPalavra = (chave, msg) => {
    const dePara = [
                        {"de":"que é","para":"quem é"},
                        {"de":"que e","para":"quem é"},
                        {"de":"quel","para":"qual"},
                        {"de":"qiel","para":"qual"},
                        {"de":"sei","para":"seu"},                        
                        {"de":"teu","para":"seu"},                        
                        {"de":"noem","para":"nome"},
                        {"de":"name","para":"nome"},
                        {"de":"nime","para":"nome"},
                        {"de":"npnw","para":"nome"},
                        {"de":"dial","para":"dia"},
                        {"de":"voce","para":"você"},
                        {"de":"vc","para":"você"},
                        {"de":"?","para":""},
                    ]
    
    if (chave.toLowerCase() === msg.toLowerCase()) {
        return true
    } 

    let msg_temp = msg.toString()

    for (let i1 = 0; i1 < dePara.length; i1++) {
        const elm1 = dePara[i1];
        msg_temp = msg_temp.replace(elm1.de, elm1.para)
        msg_temp = msg_temp.toLowerCase()
        if (chave.toLowerCase() === msg_temp.toLowerCase()) {
            return true
        } 

        if (chave.indexOf(msg_temp) > -1) {
            // verificar percentual de acerto da chave
            // 70% entao ok.
        }
    } 



    return false
}

const procFormatAcao = (elm1) => {
    var res = JSON.parse(JSON.stringify(elm1));
    if(res.Titulo.indexOf("{datetime}") > -1) {
        var dateTime = require('node-datetime');
        var dt = dateTime.create();
        var formatted = dt.format('d/m/Y H:M:S');                                    
        res.Titulo = elm1.Titulo.replace("{datetime}", " " + formatted); 
    }
    return res
}

const procSimples = (msg_recebida) =>  {
    return new Promise(      
        function(resolve, reject) {
            for (let i1 = 0; i1 < Pesquisa.length; i1++) {
                const elm1 = Pesquisa[i1];                
                for (let i2 = 0; i2 < elm1.Indexs.length; i2++) {
                    const elm2 = elm1.Indexs[i2];                    
                    if (procComparativoPalavra(elm2.chave, msg_recebida)) {                        
                        return resolve(procFormatAcao(elm1));
                    };
                }
            }
            return reject("Erro")
        }
    );
}

export const analise = (msg_recebida) => {
    return procSimples(msg_recebida)
}

