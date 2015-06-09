"use strict";

var Nuclear = require("nuclear-js");
var actionTypes = require("../action-types");

function setCategories(state, categories) {
  state = Nuclear.toImmutable({});
  for (var key in categories) {
    state = state.set(
      parseInt(key, 10),
      Nuclear.toImmutable(categories[key])
    );
  }
  return state;
}

module.exports = new Nuclear.Store({

  getInitialState() {
    return Nuclear.toImmutable({});
  },

  initialize() {
    this.on(actionTypes.SetCategories, setCategories);
  },

});

