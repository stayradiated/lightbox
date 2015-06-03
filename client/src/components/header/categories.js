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

    var items = categories.map(category => {
      return {
        label: category.get('Name'),
        onClick: this.onClick.bind(null, category),
      };
    }).toJS();

    var category = categories.find(category => {
      return category.get('ID') === activeCategory.get('ID');
    });
    
    var activeCategoryName = category ? category.get('Name') : 'Categories';

    return (
      <div className='categories'>
        <Dropdown items={items}>
          {activeCategoryName}
        </Dropdown>
      </div>
    );
  },

  onClick(category) {
    Lightbox.actions.setCategory(category.get('ID'));
  },

});

module.exports = Categories;
