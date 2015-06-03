'use strict';

var Nuclear = require('nuclear-js');
var actionTypes = require('../action-types');

module.exports = new Nuclear.Store({

  getInitialState() {
    return Nuclear.toImmutable({
      ID: -1,
    });
  },

  initialize() {
    this.on(actionTypes.SetCategory, setCategory);
  },

});

function setCategory(state, categoryID) {
  return state.set('ID', categoryID);
}
