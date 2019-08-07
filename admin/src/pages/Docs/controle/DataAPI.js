export const DataAPI =  [
    {
        "Id":1,
        "Titulo":"Cadastro de Usuários",
        "Itens":[
            {
                "Id":1,
                "Titulo":"Inserir Usuários",
                "Metodo":"POST",
                "URL":"usuario/novo",
                "Parametros":[
                    {
                        "titulo":"Código da Empresa",
                        "nome":"IDEmpresa",
                        "tipo":"text",
                        "tamanho":45,
                        "placeholder":"Código da Empresa",
                        "requerido":true,
                    },
                    {
                        "titulo":"Nome Completo",
                        "nome":"Nome",
                        "tipo":"text",
                        "tamanho":45,
                        "placeholder":"Nome/ Razão Social",
                        "requerido":true,
                    },{
                        "titulo":"E-mail",
                        "nome":"email",
                        "tipo":"email",
                        "tamanho":45,
                        "placeholder":"e-mail",
                        "requerido":true,
                    },{
                        "titulo":"CPF/ CNPJ",
                        "nome":"doc1",
                        "tipo":"text",
                        "tamanho":45,
                        "placeholder":"CPF/ CNPJ",
                        "requerido":true,
                    }
                ]
            }
        ]
    },
    {
        "Id":2,
        "Titulo":"Cadastro de Cliente",
        "Itens":[
            {
                "Titulo":"Alterar",
                "Metodo":"POST",
                "Parametros":[
                    {
                        "titulo":"Nome Completo",
                        "name":"Nome",
                        "tipo":"text",
                        "tamanho":45,
                        "placeholder":"informe o nome",
                    },{
                        "titulo":"E-mail",
                        "name":"email",
                        "tipo":"email",
                        "tamanho":45,
                        "placeholder":"informe o e-mail",
                    }
                ]
            }
        ]
    }
]
