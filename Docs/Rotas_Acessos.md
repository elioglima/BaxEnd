# Metodos - Acessos

* /api/acesso/logar/{email}/{senha}

### /api/acesso/logar/{email}/{senha}
    routes.HandleFunc("/api/acesso/logar/{email}/{senha}", use(acesso.Logar, basicAuth)).Methods("POST")
