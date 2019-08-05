import React, { Component } from 'react';
import { connect } from "react-redux";
import './css/styles.css'
import {DataAPI} from '../../controle/DataAPI'

class Objeto extends Component {
    constructor(props) {
        super(props)
        this.state = {
            registro_sel:0
        }
    }

    getDataAPI(index) {
        return DataAPI[index]
    }

    onClickRegistroSel(e, key) {
        this.setState({registro_sel:key})
    }

    render() {    
        
        return (
            <div className="docs-base"> 
                <div className="docs-menu-lateral">          
                    {
                        DataAPI.map((item, key) => {
                            return (
                                <div key={key} className="docs-menu-lateral-item" onClick={e => this.onClickRegistroSel(e, key)}>{item.Titulo}</div> 
                            )
                        })
                    }
                </div>

                <div className="docs-container">
                    { 
                        (() => {
                            const Dados = this.getDataAPI(this.state.registro_sel) 
                            return (
                                <div>
                                    <div id={Dados.Id} 
                                        className="docs-container-titulo">
                                        { Dados.Titulo } 
                                    </div>               

                                    <div className="docs-menu-lateral-itens">          
                                    {
                                            Dados.Itens.map((item, key1) => {
                                                return (
                                                    <div>
                                                        <div key={"titulo1"+key1} className="docs-menu-lateral-itens-titulo" >{item.Metodo}  {item.Titulo}</div> 

                                                        <div className="docs-menu-lateral-itens">          
                                                            {
                                                                item.Parametros.map((itemParam, key2) => {
                                                                    return (
                                                                        <div>
                                                                            <div key={"titulo-"+key2} className="docs-menu-lateral-itens-titulo" >{itemParam.tipo}</div> 
                                                                            <div key={"tipo-"+key2} className="docs-menu-lateral-itens-titulo" >{itemParam.name}</div> 
                                                                        </div>
                                                                    )
                                                                })
                                                            }
                                                        </div>   

                                                    </div>   
                                                )
                                            })
                                        }
                                    </div>   


                                </div>                     
                            )
                        })()                                                    
                    }                    
                </div>
            </div>
        )
    }
}

export default connect(null,null)(Objeto)
