import React, { Component } from 'react';
import { connect } from "react-redux";
import './index.css'

class Objeto extends Component {
    constructor(props) {
        super(props)
        
        this.state = {
            id:"CompReact"+props.nome,
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
        this.props.onChange(e, this.state.id, e.target.value)
    }

    render() {  
        return (
            <div className="CompReactTextFieldControl" >
                {
                    (() => {
                        if (this.state.titulo.length > 0) {
                            return (
                                <div className="CompReactTextFieldControlLabel" >
                                </div>
                            )
                        }                    
                    })
                }
                
                <div className="CompReactTextFieldControlInput">
                    <input  type={this.state.tipo} 
                        className={this.state.className}
                        id={this.state.id} 
                        name={this.state.id} 
                        value={this.state.vbalor} 
                        onChange={e => this.onChange(e)} 
                        placeholder={this.state.placeholder}
                        />
                </div>
                {
                    (() => {
                        if (this.state.erro.length > 0) {
                            return (
                                <div className="CompReactTextFieldControlError" >
                                </div>
                            )
                        }                    
                    })
                }
                
            </div>
        )
    }
}

export default connect(null,null)(Objeto)
