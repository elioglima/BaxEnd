import React from 'react'
import { connect } from "react-redux";
import ChatBotBase from './ChatBotBase';
import ChamadaAtendimento from './ChamadaAtendimento';
import * as processar from './Controller/Processar.js'

class Objeto extends React.Component {

    constructor(props) {
        super(props)
        this.state = {
            VisibilidadeChat: false,
            Transacoes: {
                idChat: 5001,                
                Mensagens: []
            }
        }
    }    

    componentDidMount() {
        this.props.chatbot.map((data, key) => {
            this.setState(state => {
                const obj = {msg:data.msg, identificador:"chatbot", dataexecucao:"enviado agora"}
                const list = state.Transacoes.Mensagens.push(obj);      
                return {
                  list,
                  value: obj ,
                };
              }); 
        })
    }

    onSendMensage(SendTexts) {
        if (SendTexts.length === 0) {
            return false
        }
        
        this.setState(state => {
            const obj = {msg:SendTexts, identificador:"client", dataexecucao:"enviado agora"}
            const list = state.Transacoes.Mensagens.push(obj);      
            return {
              list,
              value: obj ,
            };
        });          

        this.props.analise(SendTexts)
            .then(res => { 
                console.log("Sucesso Analise dados:", res); 
            })
            .catch(err => {
                console.log('Erro Analise dados:' + err)
            })
         
    }

    onCloseAtendimento() {
        this.setState({VisibilidadeChat: false})
    }

    onChamadaAtendimento() {
        const chatVisivel = this.state.VisibilidadeChat === false;
        this.setState({VisibilidadeChat: chatVisivel})
    }

    render() {    
        return (      
            <div>
                { 
                    (() => {
                        if (this.state.VisibilidadeChat === true) 
                            return <ChatBotBase onCloseAtendimento={this.onCloseAtendimento.bind(this)} onSendMensage={this.onSendMensage.bind(this)} Transacoes={this.state.Transacoes} />                

                        else if (this.state.VisibilidadeChat === false) 
                            return <ChamadaAtendimento onChamadaAtendimento={this.onChamadaAtendimento.bind(this)} />

                    })()
                }
            </div>      
        )
    }
}

function mapStateToProps(state) {
    return { chatbot: processar.chatbot, analise: processar.analise }
  }

export default connect(mapStateToProps, null)(Objeto)
