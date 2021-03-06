"use strict";

var React = require("react");

var flux = require("../../flux");
var Lightbox = require("../../modules/lightbox");
var ShowList = require("../common/showlist");

var TopCharts = React.createClass({
  mixins: [flux.ReactMixin],

  getDataBindings() {
    return {
      shows: Lightbox.getters.shows,
    };
  },

	render() {
    var shows = this.state.shows.sortBy(show => {
      return 0 - show.get("Rating");
    });

    return (
      <div className="route-new">
        <h1>Top Charts</h1>
        <ShowList shows={shows} />
      </div>
    );
	},

});

module.exports = TopCharts;
