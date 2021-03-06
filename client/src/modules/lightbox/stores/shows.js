"use strict";

var Nuclear = require("nuclear-js");

var actionTypes = require("../action-types");

function oldSearchFn(shows) {
  return function (query) {
    query = query.toLowerCase();
    return shows.filter(show => {
      return show.get("Title").toLowerCase().indexOf(query) > -1;
    });
  };
}

function noop() {
  return Nuclear.toImmutable([]);
}

/*
var Fuse  = require("fuse.js");

function newSearchFn(shows) {
  var fuse = new Fuse(shows.toArray(), {
    keys: ["Title"],
    getFn: fuseGetFn,
  });
  return function (query) {
    return Nuclear.toImmutable(fuse.search(query));
  };
}

function fuseGetFn(obj, path) {
  return obj.get(path);
}
*/

function setShows(state, shows) {
  shows = Nuclear.toImmutable(shows).map(show => {
    return show
      .set("Released", new Date(show.get("Released")))
      .set("DateCreated", new Date(show.get("DateCreated")));
  });

  return Nuclear.toImmutable({
    Shows: shows,
    Search: oldSearchFn(shows), // newSearchFn(shows),
  });
}

module.exports = new Nuclear.Store({

  getInitialState() {
    return Nuclear.toImmutable({
      Shows: [],
      Search: noop,
    });
  },

  initialize() {
    this.on(actionTypes.SetShows, setShows);
  },

});

