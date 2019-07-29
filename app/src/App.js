import React from 'react';
import './App.css';
import axios from 'axios';
import qs from 'qs';

class speller extends React.Component {

    constructor() {
      super();

      this.spellNum = this.spellNum.bind(this);
      this.readStr = this.readStr.bind(this);
    }

    readStr() {
      var url = 'https://api.pandyaka.com/speller/read';
      const data = qs.stringify({
        text: document.getElementById("readinput").value
      })
      console.log(document.getElementById("readinput").value);
      axios.post(url,data).then(res => document.getElementById("readoutput").innerHTML = res.data.number)
    }

    spellNum() {
      axios.get('https://api.pandyaka.com/speller/spell?number='+document.getElementById("spellinput").value).then(res => document.getElementById("spelloutput").innerHTML = res.data.text);
    }

    render() {
      return (
      <div className="App">

        <div className="App-header">
          <h1>Introducing Indonesia Numeral Spellers</h1>
        </div>

        <div className="App-container">

          <div className="spell">
            <h2>Lets spell some number..</h2>
            <form className="spellform" >
              <input type="number" className="spellinput" id="spellinput" placeholder="input me.."></input>
              <button className="spellbutton" id="spellbutton" onClick={this.spellNum}>Spell</button>
            </form>
          </div>

          <div className="spelloutput" id="spelloutput"></div>
          
          <div className="read">
            <h2>Lets read some word(s)..</h2>
            <form className="readform">
              <input type="text" className="readinput" id="readinput" placeholder="input me.."></input>
              <button className="readbutton" id="readbutton" onClick={this.readStr}>Read</button>
            </form>
          </div>

          <div className="readoutput" id="readoutput"></div>

        </div>
      </div>

      );
    }
}

export default speller;