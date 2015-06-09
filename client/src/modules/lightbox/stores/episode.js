"use strict";

var Nuclear = require("nuclear-js");
var actionTypes = require("../action-types");

function setEpisode(_, data) {
  return Nuclear.toImmutable(data).withMutations(episode => {
    return episode
      .set("DateCreated", new Date(episode.get("DateCreated")))
      .set("DatePublished", new Date(episode.get("DatePublished")))
      .set("FirstAired", new Date(episode.get("FirstAired")));
  });
}

function resetEpisiode() {
  return Nuclear.toImmutable({});
}

module.exports = new Nuclear.Store({

  getInitialState() {
    return resetEpisiode(null);
  },

  initialize() {
    this.on(actionTypes.SetEpisode, setEpisode);
    this.on(actionTypes.ResetEpisode, resetEpisiode);
  },

});

