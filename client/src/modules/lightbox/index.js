var flux = require('../../flux');

flux.registerStores({
  show:        require('./stores/show.store'),
  shows:       require('./stores/shows.store'),
  season:      require('./stores/season.store'),
  episode:     require('./stores/episode.store'),
  categories:  require('./stores/categories.store'),
  watchlist:   require('./stores/watchlist.store'),
});

module.exports = {
  actions: require('./actions'),
  getters: require('./getters')
};
