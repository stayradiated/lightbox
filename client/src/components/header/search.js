"use strict";

var React = require("react");
var { Navigation } = require("react-router");

var Search = React.createClass({
  mixins: [Navigation],

  render() {
    return (
      <form className="search" onSubmit={this.onSearch}>
        <input
          ref="search"
          type="search"
          placeholder="Search"
        />
        <button><span className="icon icon-search" /></button>
      </form>
    );
  },

  onSearch(e) {
    e.preventDefault();
    var query = this.refs.search.getDOMNode().value;
    this.transitionTo("shows", {categoryID: 0}, { query });
  },

});

module.exports = Search;
