'use strict';

var React = require('react');
var { Link } = require('react-router');

var flux     = require('../../flux');
var Lightbox = require('../../modules/lightbox');
var Dropdown = require('../common/dropdown-split-button');

var Categories = React.createClass({
  mixins: [flux.ReactMixin],

  /*
  contextTypes: {
    router: React.PropTypes.func,
  },
  */

  getDataBindings() {
    return {
      categories: Lightbox.getters.categories,
    };
  },

  render() {
    var categories = this.state.categories;
    var activeCategory = categories.find(category => {
      return category.get('ID') === 0;
    });

    if (categories == null || activeCategory == null) {
      return null;
    }

    var categoryDropdown = <Dropdown
      active={activeCategory}
      items={categories}
      linkTo='category'
      itemParams={category => {
        return {
          categoryID: category.get('ID'),
        };
      }}
      itemName={category => {
        return category.get('Name');
      }}
    />;

    return (
      <div className='categories'>
        {categoryDropdown}
      </div>
    );
  },

});

module.exports = Categories;
