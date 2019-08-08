import React, { Component } from 'react';
import { connect } from "react-redux";

class Objeto extends Component {
    constructor(props) {
        super(props)

    }

    render() {            
        return (            
            <div className={this.props.classNameBase}>                          
                { this.props.item.Id === this.props.ResponseAPIID &&  this.props.ResponseAPI.Mensagem }
                <br />
                { this.props.item.Id === this.props.ResponseAPIID &&  <div><pre>{JSON.stringify(this.props.ResponseAPI.Dados, null, 2) }</pre></div> }
            </div> 
        )
    }

}

export default connect(null,null)(Objeto)
