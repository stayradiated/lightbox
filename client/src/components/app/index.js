'use strict';

var React            = require('react');
var { RouteHandler } = require('react-router');

var Header  = require('./header');
var Sidebar = require('./sidebar');

var App = React.createClass({

  contextTypes: {
    router: React.PropTypes.func,
  },

  render() {

    var routes = this.context.router.getCurrentRoutes();
    var activeTab = "";
    if (routes.length > 1) {
      activeTab = routes[1].name;
    }

    return (
      <div className="app">
        <Header />
        <Sidebar active={activeTab} />
        <RouteHandler />
      </div>
    );
  }

});

module.exports = App;
