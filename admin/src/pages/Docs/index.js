import React, {Component} from "react";
import { connect } from "react-redux";
import * as Actions from "./actions/actions";
import '../css/styles.css';
import BannerContato from '../components/banner_contato'
import MenuOff from '../components/menu-off'
import MenuLateral from "./components/MenuLateral";

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
        <MenuOff {...this.props} label="Documentação" />
        <MenuLateral />

      </div>
      );
      }
      }

export default connect(null,Actions)(Objeto);
