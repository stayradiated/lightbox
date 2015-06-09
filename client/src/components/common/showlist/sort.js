"use strict";

var React = require("react");

var Dropdown = require("../dropdown-button");
var Lightbox = require("../../../modules/lightbox");

var Sort = React.createClass({
  propTypes: {
    sort: React.PropTypes.object.isRequired,
    onChange: React.PropTypes.func.isRequired,
  },

  render() {
    var sortBy = this.props.sort.get("By");

    var items = this.props.sort.get("List").reduce((a, v, k) => {
      a.push({
        label: k,
      });
      return a;
    }, []);

    return (
      <div className="sort">
        <Dropdown items={items} onChange={this.onChange}>
          Sort by {sortBy}
          <span className="icon icon-down-dir" />
        </Dropdown>
      </div>
    );
  },

  onChange(item) {
    Lightbox.actions.setSort(item.label);
  },

});

module.exports = Sort;
