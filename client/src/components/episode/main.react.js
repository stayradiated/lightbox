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

    var playVideo =
      "https://www.lightbox.co.nz/#/play-video/series/" + episode.get('ShowID') +
      "/season/" + episode.get('SeasonID') + 
      "/episode/" + episode.get('ID') + 
      "/media/" + episode.get('MediaID');

    return (

      <div className='route-episode'>

        <div className='poster' style={{
          backgroundImage: 'url(' + episode.get('Image') + ')'
        }}>
          <div className='overlay' />
        </div>

        <h3>{episode.get('Title')}</h3>
        <p>Episode {episode.get('Number')}</p>

        <p>{episode.get('Runtime')} minutes</p>

        <p>{episode.get('Rating')}/10 - ({episode.get('RatingCount')} votes)</p>

        <p>{episode.get('ParentalRating')} - {episode.get('ParentalRatingReason')}</p>

        <p>{episode.get('FirstAired')}</p>

        <p>{episode.get('Plot')}</p>

        <p>Directed by: {episode.get('Director')}</p>
        <p>Written by: {episode.get('Writer')}</p>
        <p>Guest Stars: {episode.get('GuestStars')}</p>

        <a target='_blank' href={'http://imdb.com/title/' + episode.get('IMDB')}>IMDB</a>
        <a target='_blank' href={playVideo}>Play Video</a>

      </div>

    );

  },

});

module.exports = Episode;
