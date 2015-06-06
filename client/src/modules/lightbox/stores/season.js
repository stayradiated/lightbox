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
  return Nuclear.toImmutable(season).update('Episodes', episodes => {
    return episodes.map(episode => {
      return episode
        .set('DateCreated', new Date(episode.get('DateCreated')))
        .set('DatePublished', new Date(episode.get('DatePublished')))
        .set('FirstAired', new Date(episode.get('FirstAired')));
    });
  });
}

function resetSeason(state) {
  return Nuclear.toImmutable({});
}
