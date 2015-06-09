"use strict";

var React = require("react");
var { Link } = require("react-router");

var flux = require("../../flux");
var Lightbox = require("../../modules/lightbox");
var Dropdown = require("../common/dropdown-split-button");

var Nav = React.createClass({
  mixins: [flux.ReactMixin],

  contextTypes: {
    router: React.PropTypes.func,
  },

  getDataBindings() {
    return {
      show: Lightbox.getters.show,
      season: Lightbox.getters.season,
      episode: Lightbox.getters.episode,
      categories: Lightbox.getters.categories,
    };
  },

  render() {
    var categoryID = parseInt(this.context.router.getCurrentParams().categoryID, 10);
    var routes = this.context.router.getCurrentRoutes();
    var route = routes[routes.length - 1].name;

    var activeShow = this.state.show;
    var activeSeason = this.state.season;
    var activeEpisode = this.state.episode;

    var showElement = null;
    var seasonElement = null;
    var episodeElement = null;
    var categoryElement = null;

    if (this.state.categories != null && this.state.categories.has(categoryID)) {
      var category = this.state.categories.get(categoryID);

      categoryElement = (
        <div className="text" key={categoryID}>
          <Link to="shows" params={{categoryID: categoryID}}>
            {category.get("Name")}
          </Link>
        </div>
      );
    }

    if (activeShow.has("ID")) {

      showElement = (
        <div className="text" key={activeShow.get("ID")}>
          <Link to="show" params={{showID: activeShow.get("ID")}}>
            {activeShow.get("Title")}
          </Link>
        </div>
      );

      if (activeSeason.has("ID")) {

        seasonElement = (
          <div className="season" key={activeSeason.get("ID")}>
            <Dropdown
              active={activeSeason}
              items={activeShow.get("Seasons")}
              linkTo="season"
              itemParams={season => {
                return {
                  showID: activeShow.get("ID"),
                  seasonID: season.get("ID"),
                };
              }}
              itemName={season => {
                return "Season " + season.get("Number");
              }}
            />
          </div>
        );

        if (activeEpisode.has("ID")) {

          episodeElement = (
            <div className="episode" key={activeEpisode.get("ID")}>
              <Dropdown
                active={activeEpisode}
                items={activeSeason.get("Episodes")}
                linkTo="episode"
                itemParams={episode => {
                  return {
                    showID: activeShow.get("ID"),
                    seasonID: activeSeason.get("ID"),
                    episodeID: episode.get("ID"),
                  };
                }}
                itemName={episode => {
                  return episode.get("Title");
                }}
              />
            </div>
          );

        }
      }
    }

    var items = [];

    switch (route) {
      case "show":
        items.push(showElement);
        break;
      case "season":
        items.push(showElement);
        items.push(seasonElement);
        break;
      case "episode":
        items.push(showElement);
        items.push(seasonElement);
        items.push(episodeElement);
        break;
      case "shows":
        items.push(categoryElement);
        break;
    }

    return (
      <div className="nav">{items}</div>
    );

  },

});

module.exports = Nav;
