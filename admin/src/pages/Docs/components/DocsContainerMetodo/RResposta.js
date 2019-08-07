import React, { Component } from 'react';
import { connect } from "react-redux";

class Objeto extends Component {
    constructor(props) {
        super(props)

    }

    render() {            
        return (            
            <div className={this.props.classNameBase}>                          
                { this.props.item.Id === this.props.ResponseAPIID &&  this.props.ResponseAPI.Response }
            </div> 
        )
    }

}

export default connect(null,null)(Objeto)
