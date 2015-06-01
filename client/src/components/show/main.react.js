'use strict';

var React = require('react');

var flux     = require('../../flux');
var Lightbox = require('../../modules/lightbox');
var Season   = require('./season.react');

var Show = React.createClass({
  mixins: [flux.ReactMixin],

  getDataBindings() {
    return {
      show: Lightbox.getters.show,
    };
  },

  componentDidMount() {
    var showID = this.props.params.showID;
    Lightbox.actions.viewShow(showID);
  },

  componentWillReceiveProps(nextProps) {
    var showID = nextProps.params.showID;
    Lightbox.actions.viewShow(showID);
  },

  render() {
    var show = this.state.show;

    if (! show.has('ID')) {
      return null;
    }

    if (!show.has('Categories')) {
      show = show.set('Categories', []);
    }

    var categories = show.get('Categories').map(category => {
      return (
        <li key={category}>{category}</li>
      );
    });

    if (!show.has('Seasons')) {
      show = show.set('Seasons', []);
    }

    var seasonElements = show.get('Seasons').map(season => {
      return (
        <Season
          key={season.get('ID')}
          season={season}
        />
      );
    });

    return (

      <div className='route-show'>

        <div className='fanart' style={{
          backgroundImage: 'url(' + show.get('Fanart') + ')'
        }} />

        <div className='poster' style={{
          backgroundImage: 'url(' + show.get('Poster') + ')'
        }}>
          <div className='overlay' />
        </div>

        <h3>{show.get('Title')}</h3>

        <p>{show.get('Runtime')}</p>

        <p>{show.get('Rating')}/10 - ({show.get('RatingCount')} votes)</p>

        <p>{show.get('Released')}</p>

        <p>{show.get('Plot')}</p>

        <p>{show.get('ParentalRating')}</p>

        <ul>
          {categories}
        </ul>

        <div className='seasons'>
          {seasonElements}
        </div>

      </div>

    );

  },

});

module.exports = Show;
