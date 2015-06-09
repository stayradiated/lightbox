"use strict";

var React = require("react");
// var TransitionGroup = require("react/lib/ReactCSSTransitionGroup");
var { RouteHandler } = require("react-router");

var Header = require("../header/");
var Sidebar = require("./sidebar");

var App = React.createClass({

  contextTypes: {
    router: React.PropTypes.func,
  },

  render() {
    // var name = this.context.router.getCurrentPath();
    var routes = this.context.router.getCurrentRoutes();
    var activeTab = "";
    if (routes.length > 1) {
      activeTab = routes[1].name;
    }

    return (
      <div className="app">
        <Header />
        <Sidebar active={activeTab} />
        <div className="route-container">
          <RouteHandler />
        </div>
      </div>
    );

    // <TransitionGroup className="route-container" component="div" transitionName="app-route">
    // </TransitionGroup>
  }

});

module.exports = App;
