'use strict';

var React = require('react');

var Sidebar = React.createClass({

  render() {
    return (
      <aside className='sidebar'>
        <h1 className='logo'>Lightbox</h1>
        <ul>
          <li>Playing</li>
          <li>Discover</li>
          <li>Activity</li>
          <li>Top Charts</li>
          <li>New Releases</li>
          <li>Watchlist</li>
        </ul>
      </aside>
    );
  },

});

module.exports = Sidebar;
