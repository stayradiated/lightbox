'use strict';

var React = require('react');

var flux       = require('../../flux');
var Lightbox   = require('../../modules/lightbox');
var Show       = require('./show.react');
var Fake       = require('./fake.react');
var Categories = require('./categories.react');
var Sort       = require('./sort.react');

var Shows = React.createClass({
  mixins: [flux.ReactMixin],

  getDataBindings() {
    return {
      shows: Lightbox.getters.shows,
    };
  },

  componentDidMount() {
    this.handleParams(this.props.params);
  },

  componentWillReceiveProps(nextProps) {
    this.handleParams(nextProps.params);
  },

  handleParams(params) {
    if (params.hasOwnProperty('query')) {
      Lightbox.actions.searchShows(params.query);
    } else if (params.hasOwnProperty('category')) {
      Lightbox.actions.viewCategory(params.category);
    } else {
      Lightbox.actions.searchShows();
    }
  },

  render() {
    var shows = this.state.shows;

    if (shows == null) {
      shows = [];
    }

    var showElements = shows.map(show => {
      return (
        <Show
          key={show.get('ID')}
          show={show}
        />
      );
    });

    // var fakeElements = new Array(total - shows.size);
    //
    // for (var i = 0, len = fakeElements.length; i < len; i += 1) {
    //   fakeElements[i] = (
    //     <Fake key={i} />
    //   );
    // }

    return (
      <div className='route-shows'>
        <header>
          <Categories />
          <Sort />
        </header>
        <div className='list'>
          {showElements}
        </div>
      </div>
    );
  },

});

module.exports = Shows;
