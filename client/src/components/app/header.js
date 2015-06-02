'use strict';

var lodash = require('lodash');
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
        <form className='search'>
          <input
            ref='search'
            type='search'
            placeholder='Search'
            onChange={lodash.debounce(this.onSearch, 300)}
          />
          <button><span className='icon-search' /></button>
        </form>
      </header>
    );
  },

  onSearch(e) {
    var query = this.refs.search.getDOMNode().value;
    this.transitionTo('search', { query });
  },

});

module.exports = Header;
