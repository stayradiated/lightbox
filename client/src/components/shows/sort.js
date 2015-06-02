'use strict';

var React = require('react');
var Router = require('react-router');
var { Link } = Router;

var Sort = React.createClass({

  render() {
    return (
      <div className='sort'>
        <p>
          Most Popular
          <span className='icon-down-dir' />
        </p>
        <ul className='dropdown'>
          <li><a>A - Z</a></li>
          <li><a>Z - A</a></li>
          <li><a>Most Popular</a></li>
        </ul>
      </div>
    )
  },

});

module.exports = Sort;
