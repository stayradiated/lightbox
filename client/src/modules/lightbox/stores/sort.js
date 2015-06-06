'use strict';

var Nuclear = require('nuclear-js');
var actionTypes = require('../action-types');

module.exports = new Nuclear.Store({

  getInitialState() {
    return Nuclear.toImmutable({
      By: 'A - Z',
      List: {
        'A - Z': sortByTitleAsc,
        'Z - A': sortByTitleDesc,
        'Popularity': sortByRatingDesc,
        'Year Released': sortByReleasedDesc,
        'Date Added': sortByDateCreatedDesc,
      }
    });
  },

  initialize() {
    this.on(actionTypes.SetSort, setSort);
  },

});

function setSort(state, sortBy) {
  return state.set('By', sortBy);
}

function sortNoop (shows) {
  return shows;
}

function sortByTitleAsc(shows) {
  return shows.sort((a, b) => {
    return a.get('Title').localeCompare(b.get('Title'));
  });
}

function sortByTitleDesc(shows) {
  return shows.sort((a, b) => {
    return b.get('Title').localeCompare(a.get('Title'));
  });
}

function sortByRatingDesc(shows) {
  return shows.sortBy(show => 0 - show.get('Rating'));
}

function sortByReleasedDesc(shows) {
  return shows.sortBy(show => 0 - show.get('Released').getTime());
}

function sortByReleasedAsc(shows) {
  return shows.sortBy(show => show.get('Released').getTime());
}

function sortByDateCreatedDesc(shows) {
  return shows.sortBy(show => 0 - show.get('ID'));
}
