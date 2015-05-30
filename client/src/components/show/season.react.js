'use strict';

var React = require('react');
var Router = require('react-router');
var { Link } = Router;

var Season = React.createClass({

  propTypes: {
    season: React.PropTypes.object,
  },

  render() {
    var season = this.props.season;

    if (! season.has('ID')) {
      return null;
    }

    console.log(season.toJS());

    return (
      <Link to='season' params={{seasonID: season.get('ID')}} className='season'>
        <img src={'http://thetvdb.com/banners/' + season.get('Banner')} />
        <h3>Season {season.get('Number')}</h3>
      </Link>
    );
  },

});

module.exports = Season;
