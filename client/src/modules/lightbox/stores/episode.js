'use strict';

var Nuclear = require('nuclear-js');
var actionTypes = require('../action-types');

module.exports = new Nuclear.Store({

  getInitialState() {
    return resetEpisiode(null);
  },

  initialize() {
    this.on(actionTypes.SetEpisode, setEpisode);
    this.on(actionTypes.ResetEpisode, resetEpisiode);
  },

});

function setEpisode(state, episode) {
  return Nuclear.toImmutable(episode);
}

function resetEpisiode(state) {
  return Nuclear.toImmutable({});
}
