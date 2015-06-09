"use strict";

var React = require("react");

var flux = require("../../../flux");
var Lightbox = require("../../../modules/lightbox");
var Show = require("../show/");
var Sort = require("./sort");

var ShowList = React.createClass({
  mixins: [flux.ReactMixin],

  propTypes: {
    shows: React.PropTypes.object,
  },

  getDataBindings() {
    return {
      sort: Lightbox.getters.sort,
    };
  },

  render() {
    var sort = this.state.sort;
    var sortFn = sort.get("List").get(sort.get("By"));
    var shows = sortFn(this.props.shows).slice(0, 50);

    var showElements = shows.map(show => {
      return (
        <Show key={show.get("ID")} show={show} />
      );
    });

    return (
      <div className="show-list">
        <Sort sort={sort} onChange={this.setSortFn} />
        <div className="elements">
          {showElements}
        </div>
      </div>
    );
  },

  setSortFn(sortFn) {
    this.setState({ sortFn });
  },

});

module.exports = ShowList;
