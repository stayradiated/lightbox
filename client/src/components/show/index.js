'use strict';

var React = require('react');
var Router = require('react-router');
var { Link } = Router;

var flux     = require('../../flux');
var Lightbox = require('../../modules/lightbox');
var Season   = require('./season');
var Rating = require('../common/rating');
var Runtime = require('../common/runtime');
var Header = require('../common/header/');
var Poster = require('../common/poster/');

var Show = React.createClass({
  mixins: [flux.ReactMixin],

  getDataBindings() {
    return {
      show: Lightbox.getters.show,
    };
  },

  render() {
    var show = this.state.show;

    if (! show.has('ID')) {
      return null;
    }

    console.log(show.toJS());

    if (!show.has('Categories')) {
      show = show.set('Categories', []);
    }

    var categories = show.get('Categories').filter(category => {
      switch (category.get('Name')) {
        case "All TV":
        case "All Kids":
          return false;
      }
      return true;
    }).map(category => {
      return (
        <li key={category.get('ID')}>
          <Link to='category' params={{categoryID: category.get('ID')}}>
            {category.get('Name')}
          </Link>
        </li>
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

        <Header show={show} />

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
                <Runtime runtime={show.get('Runtime')} />
                <span className='rating'>
                  <Rating rating={show.get('Rating')} />
                </span>
                <span className='parental-rating'>Rated {show.get('ParentalRating')}</span>
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
            <Poster url={show.get('Poster')} />
          </div>

        </div>
      </div>

    );

  },

});

module.exports = Show;
