'use strict';

var React = require('react');
var flux = require('../flux');
var Lightbox = require('../modules/lightbox');
var Series = require('./SeriesItem.react.js');

var SeriesBrowser = React.createClass({
  mixins: [flux.ReactMixin],

  getDataBindings() {
    return {
      seriesList: Lightbox.getters.seriesList,
    };
  },

  render() {
    var seriesList = this.state.seriesList;

    console.log(this.props);

    if (seriesList == null) {
      return null;
    }

    var seriesItems = seriesList.map(series => {
      return (
        <Series
          key={series.get('ID')}
          series={series}
        />
      );
    });

    return (
      <div className='series-browser'>
        <header>
          All TV
          Most Popular
        </header>
        <div className='list'>
          {seriesItems}
        </div>
      </div>
    );
  },

});

module.exports = SeriesBrowser;
