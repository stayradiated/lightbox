var flux = require('../../flux');
var actionTypes = require('./action-types');
var getters = require('./getters');

exports.showFaves = function (username) {
  flux.dispatch(actionTypes.SHOW_FAVES, {
    username: username,
  });
}
