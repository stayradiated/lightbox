'use strict';

var React            = require('react');
var { RouteHandler } = require('react-router');

var Header  = require('./header.react');
var Sidebar = require('./sidebar.react');

var App = React.createClass({

  render() {
    return (
      <div className="app">
        <Header />
        <Sidebar />
        <RouteHandler />
      </div>
    );
  }

});

module.exports = App;
