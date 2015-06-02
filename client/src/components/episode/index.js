'use strict';

var React = require('react');
var Router = require('react-router');
var moment = require('moment');
var { Link } = Router;

var flux     = require('../../flux');
var Lightbox = require('../../modules/lightbox');
var Rating = require('../common/rating');
var Header = require('../common/header/');
var Poster = require('../common/poster/');
var Runtime = require('../common/runtime');

var Episode = React.createClass({
  mixins: [flux.ReactMixin],

  getDataBindings() {
    return {
      show: Lightbox.getters.show,
      season: Lightbox.getters.season,
      episode: Lightbox.getters.episode,
    };
  },

  render() {
    var show = this.state.show;
    var season = this.state.season;
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

    var firstAired = moment(episode.get('FirstAired')).format("MMMM D, YYYY");

    return (

      <div className='route-episode'>

        <Header show={show} season={season} episode={episode} />

        <div className='contents'>

          <div className='title-container'>
            <h1>
              <Link to='show' params={{showID: show.get('ID')}}>
                {show.get('Title')}
              </Link>
            </h1>
            <h2>{episode.get('Title')}</h2>
            <h3>S{season.get('Number')} - E{episode.get('Number')}</h3>
          </div>

          <div className='metadata-container'>

            <div className='first-aired'>
              {firstAired}
            </div>

            <div className='labels'>
              <div><Runtime runtime={episode.get('Runtime')} /></div>
              <Rating rating={episode.get('Rating')} />
              <div>{episode.get('ParentalRating')} - {episode.get('ParentalRatingReason')}</div>
            </div>

            <dl>
              <dt>Director:</dt>
              <dd>{episode.get('Director')}</dd>
            </dl>

            <dl>
              <dt>Writer:</dt>
              <dd>{episode.get('Writer')}</dd>
            </dl>

            <p><a target='_blank' href={playVideo}>Play Video</a></p>

            <div className='plot'>
              <p>{episode.get('Plot')}</p>
            </div>
          </div>

          <div className='poster-container'>
            <Poster url={episode.get('Image')} />
          </div>

        </div>


      </div>

    );

  },

});

module.exports = Episode;
