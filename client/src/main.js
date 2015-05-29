var React = require('react');
var flux = require('./flux');
var Faves = require('./modules/faves');
var FavesApp = require('./components/FavesApp.react');

// export for http://fb.me/react-devtools
window.React = React;

require('./style/index.scss');

React.render(
  <FavesApp />,
  document.getElementById('react')
);
