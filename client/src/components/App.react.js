'use strict';

var React = require('react');
var Sidebar = require('./Sidebar.react');
var Browser = require('./Browser.react');

var App = React.createClass({

  render() {
    return (
      <div className="app">
        <Sidebar />
        <Browser />
      </div>
    );
  }

});

module.exports = App;
