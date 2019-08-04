import React from 'react'
import { connect } from "react-redux";
import logo from './img/icoChatBot.png';
import './css/ChamadaAtendimento.css';

/*
    base da tela
*/

class Objeto extends React.Component {

    render() {    
        return (            
            <div className="ChamadaAtendimentoBase">     
                <img src={logo} alt="Inicie seu atendimento" onClick={e=>this.props.onChamadaAtendimento(e)} />
            </div>            
        )
    }
}

export default connect(null,null)(Objeto)