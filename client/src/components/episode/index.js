"use strict";

var React = require("react");
var Router = require("react-router");
var XDate = require("xdate");
var ImmutableRenderMixin = require("react-immutable-render-mixin");
var { Link } = Router;

var flux     = require("../../flux");
var Lightbox = require("../../modules/lightbox");
var Rating = require("../common/rating");
var Poster = require("../common/poster/");
var Runtime = require("../common/runtime");

var Episode = React.createClass({
  mixins: [flux.ReactMixin, ImmutableRenderMixin],

  getDataBindings() {
    return {
      show: Lightbox.getters.show,
      season: Lightbox.getters.season,
      episode: Lightbox.getters.episode,
    };
  },

  render() {
    var show = this.state.show;
    var season = this.state.season;
    var episode = this.state.episode;

    if (!episode.has("ID")) {
      return null;
    }

    var firstAired = (new XDate(episode.get("FirstAired"))).toString("MMMM d, yyyy");

    return (

      <div className="route-episode">

        <div className="title-container">
          <h1>
            <Link to="show" params={{showID: show.get("ID")}}>
              {show.get("Title")}
            </Link>
          </h1>
          <h2>{episode.get("Title")}</h2>
          <h3>S{season.get("Number")} &#8226; E{episode.get("Number")}</h3>
        </div>

        <div className="metadata-container">

          <div className="first-aired">
            {firstAired}
          </div>

          <div className="labels">
            <span><Runtime runtime={episode.get("Runtime")} /></span>
            <span className="parental-rating" title={episode.get("ParentalRatingReason")}>{episode.get("ParentalRating")}</span>
            <Rating rating={episode.get("Rating")} />
          </div>

          { episode.has("Director") ? (
            <dl>
              <dt>Director:</dt>
              <dd>{episode.get("Director")}</dd>
            </dl>
          ) : null }

          { episode.has("Writer") ? (
            <dl>
              <dt>Writer:</dt>
              <dd>{episode.get("Writer")}</dd>
            </dl>
          ) : null }

          { episode.has("GuestStars") ? (
            <dl>
              <dt>Guest Stars:</dt>
              <dd>{episode.get("GuestStars").split(", ").slice(0, 3).join(", ")}</dd>
            </dl>
          ) : null }

          <div className="plot">
            <p>{episode.get("Plot")}</p>
          </div>
        </div>

        <div className="poster-container">
          <Poster
            id={episode.get("ID")}
            type="episodes"
            size="small"
            playOnly={true}
            onPlay={this.onPlay}
          />
        </div>

      </div>

    );
  },

  onPlay() {
    var episode = this.state.episode;

    var link =
      "https://www.lightbox.co.nz/#/play-video/series/" + episode.get("ShowID") +
      "/season/" + episode.get("SeasonID") +
      "/episode/" + episode.get("ID") +
      "/media/" + episode.get("MediaID");

    window.open(link);
  },

});

module.exports = Episode;
