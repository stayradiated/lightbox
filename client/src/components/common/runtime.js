"use strict";

var React = require("react");

var Runtime = React.createClass({

  propTypes: {
    runtime: React.PropTypes.number,
  },

  render() {
    var runtime = this.props.runtime;

    if (runtime == null) {
      return null;
    }

    var minutes = (runtime % 60);
    var hours = Math.floor(runtime / 60);
    var time = [];

    if (hours > 0) {
      time.push(hours + " hr");
    }

    if (minutes > 0) {
      time.push(minutes + " min");
    }

    return (
      <span className="runtime">{time.join(" ")}</span>
    );
  },

});

module.exports = Runtime;
