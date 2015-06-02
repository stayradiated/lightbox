'use strict';

var React = require('react');
var classNames = require('classnames');
var { Link } = require('react-router');

var Sidebar = React.createClass({

  propTypes: {
    active: React.PropTypes.string,
  },

  getDefaultProps() {
    return {
      active: 'shows',
    };
  },

  render() {

    var contents = [
      {
        name: 'Playing',
        icon: 'icon-play-circled2',
        href: 'player',
      },
      {
        name: 'Discover',
        icon: 'icon-compass',
        href: 'shows',
      },
      {
        name: 'Activity',
        icon: 'icon-history',
        href: 'activity',
      },
      {
        name: 'Top Charts',
        icon: 'icon-chart-bar',
        href: 'top',
      },
      {
        name: 'New Releases',
        icon: 'icon-calendar-empty',
        href: 'new',
      },
      {
        name: 'Watchlist',
        icon: 'icon-list',
        href: 'watchlist',
      },
    ];

    return (
      <aside className='sidebar'>
        <ul>
          {
            contents.map(tab => {
              return (
                <li key={tab.href} className={classNames({
                  active: this.props.active === tab.href,
                })}>
                  <Link to={tab.href}>
                    <span className={tab.icon} />
                    {tab.name}
                  </Link>
                </li>
              );
            })
          }
        </ul>
      </aside>
    );
  },

});

module.exports = Sidebar;
