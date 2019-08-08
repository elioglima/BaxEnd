import React, { Component } from 'react';
import { connect } from "react-redux";
import MenuLateral from '../MenuLateral';
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
                <MenuLateral {...this.props} />
                <DocsContainer {...this.props} />
            </div>
        )
    }

}

export default connect(null,null)(Objeto)
