import React, { Component } from 'react';
import { connect } from "react-redux";
import MenuAPI from '../MenuAPI';
import {DataAPI} from '../../controle/DataAPI'
import DocsContainer from '../DocsContainer';

class Objeto extends Component {
    constructor(props) {
        super(props)
    }

    getDataAPI(index) {
        return DataAPI[index]
    }

    render() {    
        return (                        
            <div className="docs-base"> 
                { this.props.registro_sel === -1 && <MenuAPI {...this.props} /> }                
                { this.props.registro_sel > -1 && <DocsContainer {...this.props} /> }                                
            </div>
        )
    }

}

export default connect(null,null)(Objeto)
