'use strict';

var Nuclear = require('nuclear-js');
var actionTypes = require('../action-types');

module.exports = new Nuclear.Store({

  getInitialState() {
    return Nuclear.toImmutable({
      seriesList: [],
    });
  },

  initialize() {
    this.on(actionTypes.SET_SERIES_LIST, setSeriesList);
  },

});

function setSeriesList(state, seriesList) {
  return state.set('seriesList', Nuclear.toImmutable(seriesList));
}
