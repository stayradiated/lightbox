'use strict';

var $        = require('jquery');
var React    = require('react');
var Router   = require('react-router');
var flux     = require('./flux');
var Lightbox = require('./modules/lightbox');

var { Route, Redirect, NotFoundRoute } = Router;

var App         = require('./components/app/');
var Show        = require('./components/show/');
var Shows       = require('./components/shows/');
var Fanart      = require('./components/fanart/');
var Season      = require('./components/season/');
var Player      = require('./components/player/');
var Episode     = require('./components/episode/');
var Activity    = require('./components/activity/');
var TopCharts   = require('./components/top/');
var WatchList   = require('./components/watchlist/');
var NewReleases = require('./components/new/');

require('./style/index.scss');

// export for http://fb.me/react-devtools
window.React = React;

var ga = require('lodash').noop;
if (window.hasOwnProperty('ga')) {
  ga = window.ga;
}

Lightbox.actions.fetchShows();
Lightbox.actions.fetchCategories();
Lightbox.actions.fetchLists();

var watchlist = localStorage.getItem('watchlist');
if (watchlist != null) {
  Lightbox.actions.setWatchlist(JSON.parse(watchlist));
}

flux.observe(
  Lightbox.getters.watchlist,
  function (items) {
    localStorage.setItem('watchlist', JSON.stringify(items.toJS()));
  }
);

var routes = (
  <Route path='/' handler={App}>

    <Redirect from='/' to='shows' params={{categoryID: 37}} />
    <NotFoundRoute handler={Shows} />

    <Route name='shows' path='shows/:categoryID' handler={Shows} />

    <Route name='fanart' handler={Fanart}>
      <Route name='show' path='/show/:showID' handler={Show} />
      <Route name='season' path='/show/:showID/season/:seasonID' handler={Season} />
      <Route name='episode' path='/show/:showID/season/:seasonID/episode/:episodeID' handler={Episode} />
    </Route>

    <Route name='watchlist' path='watchlist' handler={WatchList} />
    <Route name='new' path='new-releases' handler={NewReleases} />
    <Route name='top' path='top-charts' handler={TopCharts} />
    <Route name='player' path='playing' handler={Player} />
    <Route name='activity' path='activity' handler={Activity} />

  </Route>
);

Router.run(routes, Router.HashLocation, (Root, state) => {
  React.render(<Root />, document.getElementById('react'));

  var params = state.params;

  if (params.hasOwnProperty('showID')) {
    Lightbox.actions.fetchShow(parseInt(params.showID, 10));
  }

  if (params.hasOwnProperty('seasonID')) {
    Lightbox.actions.fetchSeason(parseInt(params.seasonID, 10));
  }

  if (params.hasOwnProperty('episodeID')) {
    Lightbox.actions.fetchEpisode(parseInt(params.episodeID, 10));
  }

  try {
    ga('send', 'pageview', {
     'page': location.pathname + location.search  + location.hash
    });
  } catch(err) {
    console.warn(err);
  }

  // if (params.hasOwnProperty('categoryID')) { 
  //   Lightbox.actions.setCategory(params.categoryID);
  // }
  //
  // if (params.hasOwnProperty('query')) { 
  //   Lightbox.actions.fetchShows(params.query);
  // }

});
