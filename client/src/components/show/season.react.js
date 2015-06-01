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

    return (
      <Link to='season' params={{seasonID: season.get('ID')}} className='season'>
        <img src={season.get('Image')} />
        <h3>Season {season.get('Number')}</h3>
      </Link>
    );
  },

});

module.exports = Season;
