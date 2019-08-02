import React, {useEffect } from 'react'
import { connect } from "react-redux";
import ChatBotBase from './ChatBotBase';
import ChamadaAtendimento from './ChamadaAtendimento';

class Objeto extends React.Component {
    withProps(Component, props) {
        return function(matchProps) {
          return <Component {...props} {...matchProps} />
        }
    }

    constructor(props) {
        super(props)
        this.state = {
            VisibilidadeChat: false,
            Transacoes: {
                idChat: 5001,                
                Mensagens: [
                    {msg:"Ola", identificador:"client", dataexecucao:"enviado à 1 minuto"},
                    {msg:"Tudo Bem", identificador:"chatbot", dataexecucao:"recebido à 1 minuto"}
                ]
            }
        }

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



export default connect(null,null)(Objeto)
