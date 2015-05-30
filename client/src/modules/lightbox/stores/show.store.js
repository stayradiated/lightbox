'use strict';

var Nuclear = require('nuclear-js');
var actionTypes = require('../action-types');

module.exports = new Nuclear.Store({

  getInitialState() {
    return Nuclear.toImmutable({});
  },

  initialize() {
    this.on(actionTypes.SetShow, setShow);
  },

});

function setShow(state, show) {
  return Nuclear.toImmutable(show);
}
