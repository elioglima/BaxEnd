import React, {Component} from "react";
import { connect } from "react-redux";
import * as Actions from "../src/actions";
import base64 from 'base-64';
import '../../css/styles.css';
import BannerContato from './components/banner-contato'
import MenuOff from './components/menu-off'

class Objeto extends Component {
  
  state = {
    name:'',
    pass:''
  }

  onSubmit = e => {
    e.preventDefault();
    var parametros = {
      N: this.state.name,
      P: this.state.pass,
    }
    this.props.Logar(parametros)    
  }  

  render() {

    this.props.Auth_app()    

    return (
      <div>  
        <BannerContato />      
        <MenuOff />
                    
        <div className="bannerss">ss
          <br />ss
        </div>
        

      </div>
      );
      }
      }

export default connect(null,Actions)(Objeto);
