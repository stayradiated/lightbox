'use strict';

var React = require('react');

var flux     = require('../../flux');
var Lightbox = require('../../modules/lightbox');

var Episode = React.createClass({
  mixins: [flux.ReactMixin],

  getDataBindings() {
    return {
      episode: Lightbox.getters.episode,
    };
  },

  componentDidMount() {
    this.handleParams(this.props.params);
  },

  componentWillReceiveProps(nextProps) {
    this.handleParams(nextProps.params);
  },

  handleParams(params) {
    Lightbox.actions.viewEpisode(params.episodeID);
  },

  render() {
    var episode = this.state.episode;

    if (! episode.has('ID')) {
      return null;
    }

    console.log(episode.toJS());

    return (

      <div className='route-episode'>

        <div className='poster' style={{
          backgroundImage: 'url(http://thetvdb.com/banners/' + episode.get('Image') + ')'
        }}>
          <div className='overlay' />
        </div>

        <h3>{episode.get('Name')}</h3>
        <h6>Episode {episode.get('Number')}</h6>

        <p>{episode.get('Runtime')}</p>

        <p>{episode.get('Rating')}/10 - ({episode.get('RatingCount')} votes)</p>

        <p>{episode.get('FirstAired')}</p>

        <p>{episode.get('Overview')}</p>

        <p>Directed by: {episode.get('Director')}</p>
        <p>Written by: {episode.get('Writer')}</p>

      </div>

    );

  },

});

module.exports = Episode;
