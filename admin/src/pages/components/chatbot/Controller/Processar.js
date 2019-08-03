
export const Pesquisa = [
        {
            "id":1,
            "Titulo":"Acesso ao Sistema",
            "Pesquisas":[
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
            "Pesquisas":[
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
            "Pesquisas":[
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

export const analise = (msg_recebida) => {
    return new Promise(      
        function(resolve, reject) {
            let retorno = {
                Response:''
            }

            let registro_localizado
            registro_localizado = false

            for (let i1 = 0; i1 < Pesquisa.length; i1++) {
                const elm1 = Pesquisa[i1];                
                for (let i2 = 0; i2 < elm1.Pesquisas.length; i2++) {
                    const elm2 = elm1.Pesquisas[i2];
                    if (elm2.chave === msg_recebida) {
                        retorno.Response = "OK";    
                        return resolve(retorno);
                    };
                }
            }

            retorno.Response = "error";    
            return reject(retorno)
        }
    );
}
    