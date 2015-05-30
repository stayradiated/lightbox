'use strict';

var React = require('react');
var Router = require('react-router');
var { Link } = Router;

var Episode = React.createClass({

  propTypes: {
    episode: React.PropTypes.object,
  },

  render() {
    var episode = this.props.episode;

    return (
      <Link to='episode' params={{episodeID: episode.get('ID')}} className='episode'>
        <img src={'http://thetvdb.com/banners/' + episode.get('Image')} />
        <h3>{episode.get('Name')}</h3>
        <p>Episode {episode.get('Number')}</p>
      </Link>
    );
  },

});

module.exports = Episode;
