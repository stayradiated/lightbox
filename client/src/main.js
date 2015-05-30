'use strict';

var $        = require('jquery');
var React    = require('react');
var Router   = require('react-router');
var flux     = require('./flux');
var Lightbox = require('./modules/lightbox');

var { Route, DefaultRoute, NotFoundRoute } = Router;

var App     = require('./components/App.react');
var Browser = require('./components/Browser.react');
var Series  = require('./components/Series.react');

require('./style/index.scss');

// export for http://fb.me/react-devtools
window.React = React;
window.Lightbox = Lightbox;

Lightbox.actions.searchSeries();

var routes = (
  <Route path='/' handler={App}>

    <DefaultRoute handler={Browser} />
    <NotFoundRoute handler={Browser} />

    <Route name='playing' path='playing' handler={Browser} />
    <Route name='series-list' path='series' handler={Browser} />
    <Route name='series' path='series/:id' handler={Series} />
    <Route name='activity' path='activity' handler={Browser} />
    <Route name='top-charts' path='top-charts' handler={Browser} />
    <Route name='new-releases' path='new-releases' handler={Browser} />
    <Route name='watchlist' path='watchlist' handler={Browser} />
  </Route>
);

Router.run(routes, Router.HashLocation, (Root, state) => {
  React.render(<Root />, document.getElementById('react'));
});
