"use strict";

var Nuclear = require("nuclear-js");
var actionTypes = require("../action-types");

function setWatchlist(state, data) {
  return Nuclear.toImmutable(data);
}

function addShowToWatchlist(state, showID) {
  var exists = state.some(item => {
    return item.get("ShowID") === showID;
  });

  if (exists) {
    return state;
  }

  return state.push(Nuclear.toImmutable({
    "Date": Date.now(),
    "ShowID": showID,
  }));
}

function removeShowFromWatchlist(state, showID) {
  return state.filter(item => {
    return item.get("ID") !== showID;
  });
}

module.exports = new Nuclear.Store({

  getInitialState() {
    return Nuclear.toImmutable([
      {
        "Date": Date.now(),
        "ShowID": 115,
      }
    ]);
  },

  initialize() {
    this.on(actionTypes.BookmarkShow, addShowToWatchlist);
    this.on(actionTypes.UnbookmarkShow, removeShowFromWatchlist);
    this.on(actionTypes.SetWatchlist, setWatchlist);
  },

});
