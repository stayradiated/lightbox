"use strict";

var React = require("react");
var Router = require("react-router");
var { Link } = Router;

var flux     = require("../../flux");
var Lightbox = require("../../modules/lightbox");
var Episode  = require("./episode");

var Season = React.createClass({
  mixins: [flux.ReactMixin],

  getDataBindings() {
    return {
      show: Lightbox.getters.show,
      season: Lightbox.getters.season,
    };
  },

  render() {
    var show = this.state.show;
    var season = this.state.season;

    if (!(show.has("ID") && season.has("ID"))) {
      return null;
    }

    var episodeElements = season.get("Episodes").map(episode => {
      return (
        <Episode
          key={episode.get("ID")}
          episode={episode}
        />
      );
    });

    return (
      <div className="route-season">

        <div className="title-container">
          <h1>
            <Link to="show" params={{showID: show.get("ID")}}>
              {show.get("Title")}
            </Link>
          </h1>
          <h2>Season {season.get("Number")}</h2>
        </div>

        <div className="metadata-container">
          {episodeElements}
        </div>

      </div>
    );
  },

});

module.exports = Season;
