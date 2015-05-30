'use strict';

var React = require('react');
var Router = require('react-router');
var { Link } = Router;

var Sidebar = React.createClass({

  render() {
    return (
      <aside className='sidebar'>
        <ul>
          <li>
            <Link to='player'>
              <span className='icon-play-circled2' />
              Playing
            </Link>
          </li>
          <li className='active'>
            <Link to='shows'>
              <span className='icon-compass' />
              Discover
            </Link>
          </li>
          <li>
            <Link to='activity'>
              <span className='icon-history' />
              Activity
            </Link>
          </li>
          <li>
            <Link to='top'>
              <span className='icon-chart-bar' />
              Top Charts
            </Link>
          </li>
          <li>
            <Link to='new'>
              <span className='icon-calendar-empty' />
              New Releases
            </Link>
          </li>
          <li>
            <Link to='watchlist'>
              <span className='icon-list' />
              Watchlist
            </Link>
          </li>
        </ul>
      </aside>
    );
  },

});

module.exports = Sidebar;
