'use strict';

var React = require('react');
var classNames = require('classnames');
var { Link } = require('react-router');

var Sidebar = React.createClass({

  propTypes: {
    active: React.PropTypes.string,
  },

  render() {

    var contents = [
      // {
      //   name: 'Playing',
      //   icon: 'icon-play-circled2',
      //   href: 'player',
      // },
      // {
      //   name: 'Activity',
      //   icon: 'icon-history',
      //   href: 'activity',
      // },
      {
        name: 'Recommended',
        icon: 'icon-thumbs-up',
        href: 'activity',
      },
      // {
      //   name: 'Top Charts',
      //   icon: 'icon-chart-bar',
      //   href: 'top',
      // },
      // {
      //   name: 'New Releases',
      //   icon: 'icon-calendar-empty',
      //   href: 'new',
      // },
      {
        name: 'Watchlist',
        icon: 'icon-list',
        href: 'watchlist',
      },
      {
        name: 'Browse',
        icon: 'icon-compass',
        href: 'shows',
        params: {categoryID: 37},
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
                  <Link to={tab.href} params={tab.params}>
                    <span className={'icon ' + tab.icon} />
                    {tab.name}
                  </Link>
                </li>
              );
            })
          }
        </ul>
        <hr />
        <ul className='categories'>
          <li>
            <Link to='shows' params={{categoryID: 37 }}>All TV</Link>
            <ul>
              <li><Link to='shows' params={{categoryID: 20 }}>Comedy</Link></li>
              <li><Link to='shows' params={{categoryID: 21 }}>Crime</Link></li>
              <li><Link to='shows' params={{categoryID: 1 }}>Drama</Link></li>
              <li><Link to='shows' params={{categoryID: 34 }}>Factual</Link></li>
              <li><Link to='shows' params={{categoryID: 39 }}>New Zealand</Link></li>
              <li><Link to='shows' params={{categoryID: 26 }}>Reality</Link></li>
              <li><Link to='shows' params={{categoryID: 38 }}>Sci-Fi / Fantasy</Link></li>
            </ul>
          </li>
          <li>
            <Link to='shows' params={{categoryID: 36 }}>Kids TV</Link>
            <ul>
              <li><Link to='shows' params={{categoryID: 35 }}>Pre-school</Link></li>
            </ul>
          </li>
        </ul>
      </aside>
    );
  },

});

module.exports = Sidebar;
