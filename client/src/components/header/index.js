"use strict";

var React = require("react");
var { Link } = require("react-router");

var Nav        = require("./nav");
var Search     = require("./search");
var Profile    = require("./profile");

var Header = React.createClass({

  render() {
    return (
      <header className="header">
        <h1 className="logo">
          <Link to="/">Lightbox</Link>
        </h1>
        <Nav />
        <Profile />
        <Search />
      </header>
    );
  },

});

module.exports = Header;
