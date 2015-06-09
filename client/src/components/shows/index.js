"use strict";

var React = require("react");

var flux     = require("../../flux");
var Lightbox = require("../../modules/lightbox");
var ShowList = require("../common/showlist");
var List     = require("../common/list");

var Shows = React.createClass({
  mixins: [flux.ReactMixin],

  getDataBindings() {
    return {
      shows: Lightbox.getters.shows,
      showsSearch: Lightbox.getters.showsSearch,
      categories: Lightbox.getters.categories,
      lists: Lightbox.getters.shows,
    };
  },

  getShows() {
    var query = this.props.query.query;
    var categoryID = this.getCategoryID();
    var shows = this.state.shows;

    if (shows == null) {
      shows = [];
    }

    // filter by query
    if (query != null) {
      return this.state.showsSearch(query);
    }

    // filter by category
    if (categoryID >= 0) {
      return shows.filter(show => {
        if (show.has("Categories")) {
          if (!show.get("Categories").contains(categoryID)) {
            return false;
          }
        }
        return true;
      });
    }

    return shows;
  },

  getCategoryID() {
    return parseInt(this.props.params.categoryID, 10);
  },

  render() {
    var shows = this.getShows();

    var category = this.state.categories.get(this.getCategoryID());

    var title = "All TV";

    var showRow = null;
    if (category != null) {
      title = category.get("Name");
      var listID = category.get("List");
      if (listID >= 0) {
        showRow = (
          <List listID={listID} />
        );
      }
    }

    return (
      <div className="route-shows">
        <h1>{title}</h1>
        { showRow ? (
          <div className="most-popular">
            <h2>Most Popular</h2>
            {showRow}
          </div>
        ) : null }
        <h2>All Shows</h2>
        <ShowList shows={shows} />
      </div>
    );
  },

});

module.exports = Shows;
