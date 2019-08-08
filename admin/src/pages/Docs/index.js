import React, {Component} from "react";
import { connect } from "react-redux";
import * as Actions from "./actions";
import './components/css/styles.css'
import BannerContato from '../components/banner_contato'
import MenuOff from '../components/menu-off'
import DocsBase from "./components/DocsBase";

class Objeto extends Component {
  
  constructor(props) {
    super(props)
    this.state = {
        registro_sel:0
    }
  }    

  onClickRegistroSel(e, key) {
      this.setState({registro_sel:key})
  }


  onSubmit = e => {
    e.preventDefault();
    var parametros = {
      N: this.state.name,
      P: this.state.pass,
    }
  }  

  render() {


    return (
      <div>  
        <BannerContato />      
        <MenuOff {...this.props} label="Documentação" />
        <DocsBase {...this.state} onClickRegistroSel={this.onClickRegistroSel.bind(this)} />
      </div>
      );
      }
      }

export default connect(null,Actions)(Objeto);
