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

    var categories = show.get('Categories').filter(category => {
      switch (category) {
        case "All TV":
        case "All Kids":
          return false;
      }
      return true;
    }).map(category => {
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

    var stars = [];
    var rating = show.get('Rating') / 2;

    var fullStars = Math.floor(rating);
    var halfStars = 0
    if (rating % 1 >= 0.75) {
      fullStars++
    } else if (rating % 1 >= 0.25) {
      halfStars++
    }
    var emptyStars = 5 - fullStars - halfStars;

    console.log(rating, fullStars, halfStars, emptyStars);

    for (var i = 0; i < fullStars; i++) {
      stars.push(
        <span key={i} className='icon-star' />
      );
    }
    for (var i = 0; i < halfStars; i++) {
      stars.push(
        <span key={fullStars+i} className='icon-star-half-alt' />
      );
    }
    for (var i = 0; i < emptyStars; i++) {
      stars.push(
        <span key={fullStars+halfStars+i} className='icon-star-empty' />
      );
    }

    return (

      <div className='route-show'>

        <header>
          Comedy > Parks and Recreation
        </header>

        <div className='contents'>

          <div className='title-container'>
            <h1>{show.get('Title')}</h1>
            <h2>{show.get('Year')}</h2>
          </div>

          <div className='metadata-container'>
            <div className='show-details'>
              <div className='categories'>
                <ul>{categories}</ul>
              </div>
              <div className='labels'>
                <span>{show.get('Runtime')} min</span>
                <span>{stars}</span>
                <span>{show.get('ParentalRating')}</span>
              </div>
              <div className='plot'>
                <p>{show.get('Plot')}</p>
              </div>
            </div>

            <div className='season-list'>
              {seasonElements}
            </div>
          </div>

          <div className='poster-container'>
            <div className='poster' style={{
              backgroundImage: 'url(' + show.get('Poster') + ')'
            }}>
              <div className='overlay' />
            </div>
          </div>

        </div>
      </div>

    );

  },

});

module.exports = Show;
