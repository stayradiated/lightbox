'use strict';

var React = require('react');

var flux     = require('../../flux');
var Lightbox = require('../../modules/lightbox');
var Episode  = require('./episode.react');

var Season = React.createClass({
  mixins: [flux.ReactMixin],

  getDataBindings() {
    return {
      season: Lightbox.getters.season,
    };
  },

  componentDidMount() {
    var seasonID = this.props.params.seasonID;
    Lightbox.actions.viewSeason(seasonID);
  },

  componentWillReceiveProps(nextProps) {
    var seasonID = nextProps.params.seasonID;
    Lightbox.actions.viewSeason(seasonID);
  },

  render() {
    var season = this.state.season;

    if (! season.has('ID')) {
      return null;
    }

    var episodeElements = season.get('Episodes').map(episode => {
      return (
        <Episode
          key={episode.get('ID')}
          episode={episode}
        />
      );
    });

    return (
      <div className='route-season'>
        <h3>Season {season.get('Number')}</h3>
        
        <div className='episodes'>
          {episodeElements}
        </div>

      </div>
    );
  },

});

module.exports = Season;
