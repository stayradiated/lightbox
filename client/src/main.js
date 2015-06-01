'use strict';

var $        = require('jquery');
var React    = require('react');
var Router   = require('react-router');
var flux     = require('./flux');
var Lightbox = require('./modules/lightbox');

var { Route, DefaultRoute, NotFoundRoute } = Router;

var App     = require('./components/app/main.react');
var Show    = require('./components/show/main.react');
var Shows   = require('./components/shows/main.react');
var Fanart  = require('./components/fanart/main.react');
var Season  = require('./components/season/main.react');
var Episode = require('./components/episode/main.react');

require('./style/index.scss');

// export for http://fb.me/react-devtools
window.React = React;
window.Lightbox = Lightbox;

Lightbox.actions.searchShows();

var routes = (
  <Route path='/' handler={App}>

    <DefaultRoute handler={Shows} />
    <NotFoundRoute handler={Shows} />

    <Route name='shows' path='shows' handler={Shows} />
    <Route name='search' path='search/:query' handler={Shows} />
    <Route name='category' path='category/:category' handler={Shows} />

    <Route name='fanart' handler={Fanart}>
      <Route name='show' path='/shows/:showID' handler={Show} />
      <Route name='season' path='/season/:seasonID' handler={Season} />
      <Route name='episode' path='/episode/:episodeID' handler={Episode} />
    </Route>

    <Route name='player' path='playing' handler={Shows} />
    <Route name='activity' path='activity' handler={Shows} />
    <Route name='top' path='top-charts' handler={Shows} />
    <Route name='new' path='new-releases' handler={Shows} />
    <Route name='watchlist' path='watchlist' handler={Shows} />
  </Route>
);

Router.run(routes, Router.HashLocation, (Root, state) => {
  React.render(<Root />, document.getElementById('react'));
});
