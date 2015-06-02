'use strict';

var React = require('react');

var Rating = React.createClass({

  propTypes: {
    rating: React.PropTypes.number.isRequired,
  },

  render() {
    var stars = [];
    var rating = this.props.rating / 2;

    var fullStars = Math.floor(rating);
    var halfStars = 0;
    if (rating % 1 >= 0.75) {
      fullStars += 1;
    } else if (rating % 1 >= 0.25) {
      halfStars += 1;
    }
    var emptyStars = 5 - fullStars - halfStars;

    console.log(rating, fullStars, halfStars, emptyStars);

    var i;
    for (i = 0; i < fullStars; i += 1) {
      stars.push(
        <span key={i} className='icon-star' />
      );
    }
    for (i = 0; i < halfStars; i += 1) {
      stars.push(
        <span key={fullStars+i} className='icon-star-half-alt' />
      );
    }
    for (i = 0; i < emptyStars; i += 1) {
      stars.push(
        <span key={fullStars+halfStars+i} className='icon-star-empty' />
      );
    }

    return (
      <div className='rating'>{stars}</div>
    );
  }

});

module.exports = Rating;
