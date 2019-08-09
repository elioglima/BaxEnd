import React, {Component} from "react";
import { connect } from "react-redux";
import * as Actions from "./actions";
import './components/css/styles.css'
import BannerContato from '../components/banner_contato'
import MenuSuperior from './components/MenuSuperior'
import DocsBase from "./components/DocsBase";

class Objeto extends Component {
  
  constructor(props) {
    super(props)
    this.state = {
        registro_sel:-1
    }
  }    

  onClickRegistroSel(e, key) {
      console.log(key)
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
        <MenuSuperior {...this.props} label="Documentação" onClickRegistroSel={this.onClickRegistroSel.bind(this)} />
        <DocsBase {...this.state} onClickRegistroSel={this.onClickRegistroSel.bind(this)} />
      </div>
      );
      }
      }

export default connect(null,Actions)(Objeto);
