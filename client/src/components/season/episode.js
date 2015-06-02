'use strict';

var React = require('react');
var Router = require('react-router');
var { Link } = Router;

var Poster = require('../common/poster/');

var Episode = React.createClass({

  propTypes: {
    episode: React.PropTypes.object,
  },

  render() {
    var episode = this.props.episode;

    return (
      <Link to='episode' params={{
        showID: episode.get('ShowID'),
        seasonID: episode.get('SeasonID'),
        episodeID: episode.get('ID'),
      }} className='episode'>
        <Poster url={episode.get('Image')} />
        <h3>{episode.get('Title')}</h3>
        <p>Episode {episode.get('Number')}</p>
      </Link>
    );
  },

});

module.exports = Episode;
