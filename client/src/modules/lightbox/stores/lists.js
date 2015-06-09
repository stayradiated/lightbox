"use strict";

var Nuclear = require("nuclear-js");
var actionTypes = require("../action-types");

function setLists(state, lists) {
  return Nuclear.toImmutable(lists);
}

module.exports = new Nuclear.Store({

  getInitialState() {
    return Nuclear.toImmutable([]);
  },

  initialize() {
    this.on(actionTypes.SetLists, setLists);
  },

});
