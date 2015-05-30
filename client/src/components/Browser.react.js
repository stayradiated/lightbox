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
        {seriesItems}
      </div>
    );
  },

});

module.exports = SeriesBrowser;
