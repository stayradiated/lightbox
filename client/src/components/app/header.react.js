'use strict';

var React = require('react');
var Router = require('react-router');
var { Link, Navigation } = Router;

var Lightbox = require('../../modules/lightbox');

var Header = React.createClass({
  mixins: [Navigation],

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
        <div className='user-settings'>
          <button type='button' className='name'>
            George
            <span className='icon-down-open' />
          </button>
          <ul className='dropdown'>
            <li><a href='#'>Add Profile</a></li>
            <li><a href='#'>Settings</a></li>
            <li><a href='#'>Help</a></li>
            <li><a href='#'>Logout</a></li>
          </ul>
        </div>
      </header>
    );
  },

  onSearch(e) {
    e.preventDefault();
    var query = this.refs.search.getDOMNode().value;
    this.transitionTo('search', { query });
  },

});

module.exports = Header;
