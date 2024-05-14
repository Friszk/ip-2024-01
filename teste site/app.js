'use strict'

const switcher = document.querySelector('.btn');

switcher.addEventListener('click', function() {
  document.body.classList.toggle('black-theme')
  
  var className = document.body.className;
  if (className == "light-theme") {
    this.textContent = "Black";
  } 
  else {
    this.textContent = "Light";
  }
  
});