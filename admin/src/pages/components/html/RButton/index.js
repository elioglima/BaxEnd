import React, { Component } from 'react';
import { connect } from "react-redux";
import './index.css'

class Objeto extends Component {
    constructor(props) {
        super(props)
        
        this.state = {
            id:"CompReact"+props.nome+props.tipo,
            nome:props.nome,
            tipo:props.tipo,
            valor:props.valor,
            placeholder:props.placeholder,
            titulo:props.titulo,
            erro:props.erro,
            className:props.className,
        }

    }

    onChange = (e) => {
        this.setState({valor:e.target.value})
    }

    render() {  
        return (
            <div className="CompReactRButtonControl" >
                {
                    (() => {
                        if (this.state.titulo.length > 0) {
                            return (
                                <div className="CompReactRButtonControlLabel" >
                                </div>
                            )
                        }                    
                    })
                }
                
                <div className="CompReactRButtonControlInput">
                    <button  type={this.state.tipo} 
                        className={this.state.className}
                        id={this.state.id} 
                        name={this.state.id}                         
                        onClick={e => this.props.onClick(this.props.item, e)} 
                        >{this.props.valor}</button> 
                </div>
                
            </div>
        )
    }
}

export default connect(null,null)(Objeto)
