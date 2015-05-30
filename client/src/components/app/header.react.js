'use strict';

var React = require('react');
var Router = require('react-router');
var { Link } = Router;

var Lightbox = require('../../modules/lightbox');

var Header = React.createClass({

  render() {
    return (
      <header className='header'>
        <h1 className='logo'>
          <Link to='/'>Lightbox</Link>
        </h1>
        <h2>Discover</h2>
        <form onSubmit={this.onSearch}>
          <input ref='search' type='search' placeholder='Search' />
          <button>Search</button>
        </form>
        <p>George</p>
      </header>
    );
  },

  onSearch(e) {
    e.preventDefault();
    Lightbox.actions.searchShows(this.refs.search.getDOMNode().value);
  },

});

module.exports = Header;
