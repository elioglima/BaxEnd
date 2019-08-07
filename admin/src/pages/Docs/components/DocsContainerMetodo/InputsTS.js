import React, { Component } from 'react';
import { connect } from "react-redux";
import TextField from '../../../components/html/TextField';

class Objeto extends Component {
    constructor(props) {
        super(props)
    }
    
    


    render() {    
        return (            
            <div className="docs-container-metodo-input">          
                {                                         
                    this.props.item.Parametros.map((itemParam, key2) => {
                        let nome = "ID" + this.props.item.Id + "I" + key2 + "NMCP" + itemParam.nome                        
                        return (
                            <div>  
                                <TextField 
                                    nome={nome} 
                                    tipo={itemParam.tipo} 
                                    placeholder={itemParam.placeholder}  
                                    onChange={this.props.onChange}  
                                    />                                                                                                                              
                            </div>
                        )
                    })
                }
            </div> 
        )
    }

}

export default connect(null,null)(Objeto)
