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
  return Nuclear.toImmutable(episode).withMutations(episode => {
    return episode
      .set('DateCreated', new Date(episode.get('DateCreated')))
      .set('DatePublished', new Date(episode.get('DatePublished')))
      .set('FirstAired', new Date(episode.get('FirstAired')));
  });
}

function resetEpisiode(state) {
  return Nuclear.toImmutable({});
}
