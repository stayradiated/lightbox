'use strict';

var React = require('react');
var { Link } = require('react-router');

var flux     = require('../../flux');
var Lightbox = require('../../modules/lightbox');
var Dropdown = require('../common/dropdown-button');

var Categories = React.createClass({
  mixins: [flux.ReactMixin],

  getDataBindings() {
    return {
      category: Lightbox.getters.category,
      categories: Lightbox.getters.categories,
    };
  },

  render() {
    var activeCategory = this.state.category;
    var categories = this.state.categories;

    if (categories == null || activeCategory == null) {
      return null;
    }

    var items = categories.map((name, id) => {
      return {
        label: name,
        onClick: this.onClick.bind(null, id),
      };
    }).toList().toJS();

    var activeCategoryName = categories.get(activeCategory.get('ID')) || 'Categories';

    return (
      <div className='categories'>
      </div>
    );
  },

  onClick(categoryID) {
    Lightbox.actions.setCategory(categoryID);
  },

});

module.exports = Categories;
