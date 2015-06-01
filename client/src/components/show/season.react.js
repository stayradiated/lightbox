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

    var style = {
      backgroundImage: 'url(' + season.get('Image') + ')',
    };

    return (
      <Link to='season' params={{seasonID: season.get('ID')}} className='season'>
        <div className='poster' style={style}>
          <div className='overlay' />
        </div>
        <h3>Season {season.get('Number')}</h3>
        <p>{season.get('EpisodeCount')} Episodes</p>
      </Link>
    );
  },

});

module.exports = Season;
