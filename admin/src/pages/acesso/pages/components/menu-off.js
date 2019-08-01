import React, { Component } from 'react';
import { connect } from "react-redux";
import '../../css/styles.css'

class Objeto extends Component {
  render() {    
    return (
      <div> 
        <div className="menu-off">
          <div className="col-30">
            <span className="titulo">Baxend</span>            
          </div>

          <div className="col-70 base-right">
            <span className="item">Inicio</span>
            <span className="item">Documetação</span>
            <span className="item">Suporte</span>
          </div>
        </div>
        <div className="divisao"></div>
      </div>
    )
  }
}

export default connect(null,null)(Objeto)
