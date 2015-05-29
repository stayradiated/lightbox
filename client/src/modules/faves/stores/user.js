var Nuclear = require('nuclear-js');
var toImmutable = Nuclear.toImmutable;
var actionTypes = require('../action-types');

module.exports = new Nuclear.Store({

  getInitialState() {
    return toImmutable({})
  },

  initialize() {
    this.on(actionTypes.SHOW_FAVES, showFaves)
  },

});

/**
 * Show favorites for pandora user name
 * @param {Immutable.Map}
 * @param {Object} payload
 */
function showFaves(state, { username }) {
  return state.set('username', username);
}
