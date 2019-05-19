import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';

ReactDOM.render(<App />, document.getElementById('root'));

document.getElementById("spellbutton").addEventListener("click", function(e){
    e.preventDefault();
  })

document.getElementById("readbutton").addEventListener("click", function(e){
    e.preventDefault();
  })