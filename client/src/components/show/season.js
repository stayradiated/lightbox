"use strict";

var React = require("react");
var Router = require("react-router");
var { Link } = Router;

var Poster = require("../common/poster/");

var Season = React.createClass({

  propTypes: {
    season: React.PropTypes.object,
  },

  render() {
    var season = this.props.season;

    if (!season.has("ID")) {
      return null;
    }

    return (
      <Link to="season" params={{
        showID: season.get("ShowID"),
        seasonID: season.get("ID"),
      }} className="season">
        <Poster id={season.get("ID")} type="seasons" />
        <h4>Season {season.get("Number")}</h4>
        <h5>{season.get("EpisodeCount")} Episodes</h5>
      </Link>
    );
  },

});

module.exports = Season;
