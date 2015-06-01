'use strict';

var Nuclear = require('nuclear-js');
var actionTypes = require('../action-types');

module.exports = new Nuclear.Store({

  getInitialState() {
    return Nuclear.toImmutable({
      Data: [],
      Total: 0,
    });
  },

  initialize() {
    this.on(actionTypes.SetShows, setShows);
  },

});

function setShows(state, shows) {
  return Nuclear.toImmutable(shows);
}
