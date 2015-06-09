var flux = require("../../flux");

flux.registerStores({
  show:        require("./stores/show"),
  sort:        require("./stores/sort"),
  shows:       require("./stores/shows"),
  lists:       require("./stores/lists"),
  season:      require("./stores/season"),
  episode:     require("./stores/episode"),
  categories:  require("./stores/categories"),
  watchlist:   require("./stores/watchlist"),
});

module.exports = {
  actions: require("./actions"),
  getters: require("./getters")
};
