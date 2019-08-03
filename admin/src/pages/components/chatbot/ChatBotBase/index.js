import React from 'react'
import { connect } from "react-redux";
import './css/styles.css'
import './css/ChatBotText.css';
import ChatBotHeader from '../ChatBotHeader';
import ChatBotSendMsg from '../ChatBotSendMsg';
import ScrollToBottom from 'react-scroll-to-bottom';

/*
    base da tela
*/

class Objeto extends React.Component {
    constructor(props) {
        super(props)   
        this.state = {
            Transacoes: props.Transacoes,
        }     
    }

    
    getChatBotTextMsgColor(identificador) {
        if (identificador === 'chatbot') {
            return "ChatBotTextMsgServer"        
        } else if (identificador === 'client') {
            return "ChatBotTextMsgClient"
        }
    }

    getChatBotTituloMsgColor(identificador) {
        if (identificador === 'chatbot') {
            return "ChatBotTextIdentificadorServer"        
        } else if (identificador === 'client') {
            return "ChatBotTextIdentificadorClient"
        }
    }

    getChatBotIdentificadorMsg(item) {
        if (item.identificador === 'chatbot') {
            return "atendimento - " + item.dataexecucao        
        } else if (item.identificador === 'client') {
            return item.dataexecucao + " - eu"        
        }
    }


    render() {            
        return (            

            <div className="ChatBotBase">                
                <ChatBotHeader onCloseAtendimento={this.props.onCloseAtendimento} />

                <ScrollToBottom className="ChatBotText">                 
                    {
                        this.state.Transacoes.Mensagens.map((item, key) => {
                               return ( 
                                    <div className="ChatBotTextBase" key={key}  >
                                        <div className={this.getChatBotTituloMsgColor(item.identificador)}>{this.getChatBotIdentificadorMsg(item)}</div>
                                        <div className={this.getChatBotTextMsgColor(item.identificador)}>{item.msg}</div>
                                    </div>
                               )
                            }   
                        )
                    }                
                </ScrollToBottom>      

                <ChatBotSendMsg onSendMensage={this.props.onSendMensage}  />
            </div>            
        )
    }
}

export default connect(null,null)(Objeto)