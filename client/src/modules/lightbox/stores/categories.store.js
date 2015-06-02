'use strict';

var Nuclear = require('nuclear-js');
var actionTypes = require('../action-types');

module.exports = new Nuclear.Store({

  getInitialState() {
    return Nuclear.toImmutable([]);
  },

  initialize() {
    this.on(actionTypes.SetCategories, setCategories);
  },

});

function setCategories(state, categories) {
  return Nuclear.toImmutable(categories);
}
