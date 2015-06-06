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
    var categories = this.state.categories;

    if (! show.has('ID')) {
      return null;
    }

    if (!show.has('Categories')) {
      show = show.set('Categories', []);
    }

    var categoryList = show.get('Categories').filter(categoryID => {
      switch (categoryID) {
        case 0:
        case 37:
        case 36:
          return false;
      }
      return true;
    }).map(categoryID => {

      var category = categories.get(categoryID);

      return (
        <li key={categoryID}>
          <Link to='shows' params={{categoryID: categoryID}}>
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
          <h3>{show.get('Released').getFullYear()}</h3>
        </div>

        <div className='metadata-container'>
          <div className='show-details'>
            <div className='categories'>
              <ul>{categoryList}</ul>
            </div>
            <div className='labels'>
              <Runtime runtime={show.get('Runtime')} />
              <span className='parental-rating'>{show.get('ParentalRating')}</span>
              <Rating rating={show.get('Rating')} />
            </div>

            { show.has('Writer') ? (
              <dl>
                <dt>Creators:</dt>
                <dd>{show.get('Writer')}</dd>
              </dl>
            ) : null }


            { show.has('Actors') ? (
              <dl>
                <dt>Stars:</dt>
                <dd>{show.get('Actors').split(", ").slice(0,3).join(", ")}</dd>
              </dl>
            ) : null }

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
