import React, { Component } from 'react';
import { connect } from "react-redux";
import RButton from '../../../components/html/RButton';
import InputsTS from './InputsTS.js'
import * as Actions from '../../actions'
import RResposta from './RResposta';

class Objeto extends Component {
    
    constructor(props) {
        super(props)
        this.state = {
            "ResponseAPI":{
                "Status":0,
                "Response":""
            },
            "ResponseAPIID":0
        }
    }

    onChange(e, nome, valor) {        
        this.setState({[nome]:valor})
    }
   
    onClick(item, e) {

        let DadosJson = {}
        for (const key in this.state) {
            if (this.state.hasOwnProperty(key)) {
                if (key.indexOf("CompReactID"+item.Id) > -1) {                     
                    const nome = this.props.getName(key)
                    DadosJson[nome] = this.state[key]
                } 
            }
        }


        // [{nome:valor},{nome2:valor}]
        // {nome:valor, nome2:valor}
        
        var count = Object.keys(DadosJson).length;
        if (count === 0) {
            return
        }

        console.log(count, DadosJson);

        this.props.dispRAPI(item.URL, DadosJson)
            .then(res => {
                console.log("sucesso", res)
                // this.setState(
                //     {
                //         "ResponseAPI":{
                //             Status:res.Status,
                //             Response:res.body.Msg,
                //         },
                //         "ResponseAPIID":item.Id
                //     })

            })
            .catch(erro => {
                // console.log("erro", erro.Status, erro.Response.message)

                this.setState(
                    {
                        "ResponseAPI":{
                            Status:erro.Status,
                            Response:erro.Response.message,
                        },
                        "ResponseAPIID":item.Id
                    })

                // console.log('erro', erro, item.Id)                
            })

        // console.log("DadosForm", DadosForm)    

    }

    render() {           
        
        return (            
            <div>
                {
                    this.props.Dados.Itens.map((item, key1) => {     

                        return (
                            <div className="docs-container-metodo-base">          

                                <div key={"titulo1"+key1} className="docs-container-metodo-titulo" >
                                    <span className={item.Metodo}>{item.Metodo}</span>
                                    <span>{item.Titulo}</span>
                                </div> 

                                <div key={"url"+key1} className="docs-container-metodo-url" >{item.URL}</div> 
                                
                                <InputsTS item={item} onChange={this.onChange.bind(this)} />

                                <div className="docs-container-metodo-botoes">          
                                    <RButton 
                                        item={item}
                                        className={"Executar"} 
                                        valor="Executar" 
                                        onClick={this.onClick.bind(this)}
                                        />
                                </div>

                                <RResposta item={item} ResponseAPIID={this.state.ResponseAPIID} ResponseAPI={this.state.ResponseAPI} />
                            </div>   
                        )
                    })
                }
            </div>   
        )
    }

}


export default connect(null,Actions)(Objeto)
