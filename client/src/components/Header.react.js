'use strict';

var React = require('react');
var Lightbox = require('../modules/lightbox');

var Header = React.createClass({

  render() {
    return (
      <header className='header'>
        <h1 className='logo'>Lightbox</h1>
        <h2>Discover</h2>
        <form onSubmit={this.onSearch}>
          <input ref='search' type='search' placeholder='Search' onChang/>
          <button>Search</button>
        </form>
        <p>George</p>
      </header>
    );
  },

  onSearch(e) {
    e.preventDefault();
    Lightbox.actions.searchSeries(this.refs.search.getDOMNode().value);
  },

});

module.exports = Header;
