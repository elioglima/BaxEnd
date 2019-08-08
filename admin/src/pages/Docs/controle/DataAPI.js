export const DataAPI =  [
    {
        "Id":1,
        "Titulo":"Cadastro de Usuários",
        "Itens":[
            {
                "Id":1,
                "Titulo":"Pesquisa Todos Usuários",
                "Metodo":"POST",
                "URL":"usuario/pesquisa/todos",
                "Parametros":[
                    {
                        "titulo":"Código da Empresa",
                        "nome":"EmpresaID",
                        "tipo":"number",
                        "tamanho":45,
                        "placeholder":"Código da Empresa",
                        "requerido":true,
                        "valor":1
                    }
                ]
            }
        ]
    }
]
