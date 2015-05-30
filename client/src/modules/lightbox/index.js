var flux = require('../../flux');

flux.registerStores({
  discover: require('./stores/discover-store'),
});

module.exports = {
  actions: require('./actions'),
  getters: require('./getters')
};
