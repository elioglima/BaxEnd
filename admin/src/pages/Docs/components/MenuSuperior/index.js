import React, { Component } from 'react';
import { connect } from "react-redux";
import * as Actions from '../../../../routes/routes_actions'
import './css/styles.css';

class Objeto extends Component {
  render() {    
    return (
      <div> 
        <div className="menu-off">
          <div className="col-30">
            <span className="titulo">{this.props.label}</span>            
          </div>

          <div className="col-70 base-right">
            <span className="item" onClick={e => this.props.dispHome(e)}>Inicio</span>
            <span className="item" onClick={e => this.props.onClickRegistroSel(e, -1)}>Pesquisa</span>
            <span className="item">Suporte</span> 
          </div>
        </div>
        <div className="divisao"></div>
      </div>
    )
  }
}

export default connect(null,Actions)(Objeto)
