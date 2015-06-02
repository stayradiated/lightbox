'use strict';

var $        = require('jquery');
var React    = require('react');
var Router   = require('react-router');
var flux     = require('./flux');
var Lightbox = require('./modules/lightbox');

var { Route, DefaultRoute, NotFoundRoute } = Router;

var App       = require('./components/app/');
var Show      = require('./components/show/');
var Shows     = require('./components/shows/');
var Fanart    = require('./components/fanart/');
var Season    = require('./components/season/');
var Episode   = require('./components/episode/');
var WatchList = require('./components/watchlist/');

require('./style/index.scss');

// export for http://fb.me/react-devtools
window.React = React;
window.Lightbox = Lightbox;

Lightbox.actions.fetchShows();
Lightbox.actions.fetchCategories();

flux.observe(
  Lightbox.getters.watchlist,
  function (items) {
    localStorage.setItem('watchlist', JSON.stringify(items.toJS()));
  }
);

var routes = (
  <Route path='/' handler={App}>

    <DefaultRoute handler={Shows} />
    <NotFoundRoute handler={Shows} />

    <Route name='shows' path='shows' handler={Shows} />
    <Route name='search' path='search/:query' handler={Shows} />
    <Route name='category' path='category/:categoryID' handler={Shows} />

    <Route name='watchlist' path='watchlist' handler={WatchList} />

    <Route name='fanart' handler={Fanart}>
      <Route name='show' path='/show/:showID' handler={Show} />
      <Route name='season' path='/show/:showID/season/:seasonID' handler={Season} />
      <Route name='episode' path='/show/:showID/season/:seasonID/episode/:episodeID' handler={Episode} />
    </Route>

    <Route name='player' path='playing' handler={Shows} />
    <Route name='activity' path='activity' handler={Shows} />
    <Route name='top' path='top-charts' handler={Shows} />
    <Route name='new' path='new-releases' handler={Shows} />
  </Route>
);

Router.run(routes, Router.HashLocation, (Root, state) => {

  var params = state.params;

  if (params.hasOwnProperty('showID')) {
    Lightbox.actions.fetchShow(params.showID);
  }

  if (params.hasOwnProperty('seasonID')) {
    Lightbox.actions.fetchSeason(params.seasonID);
  }

  if (params.hasOwnProperty('episodeID')) {
    Lightbox.actions.fetchEpisode(params.episodeID);
  }

  if (params.hasOwnProperty('query')) { 
    Lightbox.actions.fetchShows(params.query);
  }

  if (params.hasOwnProperty('categoryID')) {
    Lightbox.actions.fetchCategories(params.categoryID);
  }

  React.render(<Root />, document.getElementById('react'));
});
