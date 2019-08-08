import React, { Component } from 'react';
import { connect } from "react-redux";
import {DataAPI} from '../../controle/DataAPI'
import DocsContainer from '../DocsContainer';

class Objeto extends Component {
    constructor(props) {
        super(props)
    }   
    
    render() {    
        return (
            <div className="docs-menu-lateral">          
                {
                    DataAPI.map((item, key) => {
                        return (
                            <div key={key} className="docs-menu-lateral-item" onClick={e => this.props.onClickRegistroSel(e, key)}>{item.Titulo}</div> 
                        )
                    })
                }
            </div>
        )
    }
}

export default connect(null,null)(Objeto)
