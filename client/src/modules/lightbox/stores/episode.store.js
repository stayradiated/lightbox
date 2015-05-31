'use strict';

var Nuclear = require('nuclear-js');
var actionTypes = require('../action-types');

module.exports = new Nuclear.Store({

  getInitialState() {
    return Nuclear.toImmutable({});
  },

  initialize() {
    this.on(actionTypes.SetEpisode, setEpisode);
  },

});

function setEpisode(state, episode) {
  return Nuclear.toImmutable(episode);
}
