import React, { Component } from 'react';
import { connect } from "react-redux";
import base64 from 'base-64';

class Objeto extends Component {

  state = {
    name:'',
    pass:''
  }

  onSubmit = e => {
    e.preventDefault();
    var parametros = {
      N: this.state.name,
      P: this.state.pass,
    }
    this.props.Logar(parametros)    
  } 

  render() {    
    return (
      <div> 
        <div className="banner-contato">
          <div className="col-50">
            <i class="fas fa-phone-alt"></i>
            <span className="titulo">Telefone</span>
            <span className="texto">(11) 2082-8568 / Whats: (11) 95255-0331</span>
          </div>

          <div className="col-50">
            <div className="col-50">
              <i class="far fa-envelope"></i>
              <span className="titulo">Email</span>
              <span className="texto">contato@baxend.com.br</span>
            </div>
            
          </div>

        </div>
      </div>
    )
  }
}

export default connect(null,null)(Objeto)
