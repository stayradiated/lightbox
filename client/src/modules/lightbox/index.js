var flux = require('../../flux');

flux.registerStores({
  show:        require('./stores/show'),
  shows:       require('./stores/shows'),
  season:      require('./stores/season'),
  episode:     require('./stores/episode'),
  category:    require('./stores/category'),
  categories:  require('./stores/categories'),
  watchlist:   require('./stores/watchlist'),
});

module.exports = {
  actions: require('./actions'),
  getters: require('./getters')
};
