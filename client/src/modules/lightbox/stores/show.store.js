'use strict';

var Nuclear = require('nuclear-js');
var actionTypes = require('../action-types');

module.exports = new Nuclear.Store({

  getInitialState() {
    return resetShow(null);
  },

  initialize() {
    this.on(actionTypes.SetShow, setShow);
    this.on(actionTypes.ResetShow, resetShow);
  },

});

function setShow(state, show) {
  return Nuclear.toImmutable(show);
}

function resetShow(state) {
  return Nuclear.toImmutable({});
}
