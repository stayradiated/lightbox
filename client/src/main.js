'use strict';

var $ = require('jquery');
var React = require('react');
var flux = require('./flux');
var Lightbox = require('./modules/lightbox');
var App = require('./components/App.react');

require('./style/index.scss');

// export for http://fb.me/react-devtools
window.React = React;

Lightbox.actions.searchSeries();

React.render(
  <App />,
  document.getElementById('react')
);
