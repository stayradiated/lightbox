'use strict';

var lodash = require('lodash');
var React = require('react');
var { Navigation } = require('react-router');

var Search = React.createClass({
  mixins: [Navigation],

  render() {
    return (
      <div className='search'>
        <input
          ref='search'
          type='search'
          placeholder='Search'
          onChange={lodash.debounce(this.onSearch, 300)}
        />
        <span className='icon icon-search' />
      </div>
    );
  },

  onSearch(e) {
    var query = this.refs.search.getDOMNode().value;
    this.transitionTo('search', { query });
  },

});

module.exports = Search;
