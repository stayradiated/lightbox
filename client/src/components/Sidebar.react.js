'use strict';

var React = require('react');

var Sidebar = React.createClass({

  render() {
    return (
      <aside className='sidebar'>
        <ul>
          <li>
            <a href='#/playing'>
              <span className='icon-play-circled2' />
              Playing
            </a>
          </li>
          <li className='active'>
            <a href='#/series'>
              <span className='icon-compass' />
              Discover
            </a>
          </li>
          <li>
            <a href='#/activity'>
              <span className='icon-history' />
              Activity
            </a>
          </li>
          <li>
            <a href='#/top-charts'>
              <span className='icon-chart-bar' />
              Top Charts
            </a>
          </li>
          <li>
            <a href='#/new-releases'>
              <span className='icon-calendar-empty' />
              New Releases
            </a>
          </li>
          <li>
            <a href='#/watchlist'>
              <span className='icon-list' />
              Watchlist
            </a>
          </li>
        </ul>
      </aside>
    );
  },

});

module.exports = Sidebar;
