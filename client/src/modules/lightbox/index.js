var flux = require('../../flux');

flux.registerStores({
  show:    require('./stores/show.store'),
  shows:   require('./stores/shows.store'),
  season:  require('./stores/season.store'),
});

module.exports = {
  actions: require('./actions'),
  getters: require('./getters')
};
