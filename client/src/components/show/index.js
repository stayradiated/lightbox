'use strict';

var React = require('react');
var Router = require('react-router');
var { Link } = Router;

var flux     = require('../../flux');
var Lightbox = require('../../modules/lightbox');
var Season   = require('./season');
var Rating = require('../common/rating');
var Runtime = require('../common/runtime');
var Poster = require('../common/poster/');

var Show = React.createClass({
  mixins: [flux.ReactMixin],

  getDataBindings() {
    return {
      show: Lightbox.getters.show,
      categories: Lightbox.getters.categories,
    };
  },

  render() {
    var show = this.state.show;
    var categoryList = this.state.categories;

    if (! show.has('ID')) {
      return null;
    }

    if (!show.has('Categories')) {
      show = show.set('Categories', []);
    }

    var categories = show.get('Categories').map(categoryID => {
      var category = categoryList.find(c => {
        c.get('ID') === categoryID;
      });
      if (category == null) {
        return null;
      }
      return (
        <li key={categoryID}>
          <Link to='category' params={{categoryID: categoryID}}>
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
          <Poster id={show.get('ID')} type='shows' size='large' />
        </div>

      </div>
    );
  },
});

module.exports = Show;
