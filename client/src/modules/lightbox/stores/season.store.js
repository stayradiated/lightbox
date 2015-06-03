'use strict';

var Nuclear = require('nuclear-js');
var actionTypes = require('../action-types');

module.exports = new Nuclear.Store({

  getInitialState() {
    return resetSeason();
  },

  initialize() {
    this.on(actionTypes.SetSeason, setSeason);
    this.on(actionTypes.ResetSeason, resetSeason);
  },

});

function setSeason(state, season) {
  return Nuclear.toImmutable(season);
}

function resetSeason(state) {
  return Nuclear.toImmutable({});
}
