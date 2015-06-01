'use strict';

var React = require('react');
var Router = require('react-router');
var { Link } = Router;

var Categories = React.createClass({

  render() {
    return (
      <div className='categories'>
        <p>
          All TV
          <span className='icon-down-dir' />
        </p>
        <ul className='dropdown'>
          <li>
            <Link to='category' params={{category: 0}}>
              New Zealand
            </Link>
          </li>
          <li>
            <Link to='category' params={{category: 1}}>
              Drama
            </Link>
          </li>
          <li>
            <Link to='category' params={{category: 2}}>
              Mystery
            </Link>
          </li>
          <li>
            <Link to='category' params={{category: 3}}>
              Horror
            </Link>
          </li>
          <li>
            <Link to='category' params={{category: 4}}>
              Family
            </Link>
          </li>
          <li>
            <Link to='category' params={{category: 5}}>
              Short
            </Link>
          </li>
          <li>
            <Link to='category' params={{category: 6}}>
              Sci-Fi
            </Link>
          </li>
          <li>
            <Link to='category' params={{category: 7}}>
              Reality-TV
            </Link>
          </li>
          <li>
            <Link to='category' params={{category: 8}}>
              War
            </Link>
          </li>
          <li>
            <Link to='category' params={{category: 9}}>
              Biography
            </Link>
          </li>
            <li>
            <Link to='category' params={{category: 10}}>
              History
            </Link>
          </li>
            <li>
            <Link to='category' params={{category: 11}}>
              Romance
            </Link>
          </li>
          <li>
            <Link to='category' params={{category: 12}}>
              Musical
            </Link>
          </li>
            <li>
            <Link to='category' params={{category: 13}}>
              Talk-Show
            </Link>
          </li>
          <li>
            <Link to='category' params={{category: 14}}>
              Music
            </Link>
          </li>
          <li>
            <Link to='category' params={{category: 15}}>
              Sport
            </Link>
          </li>
          <li>
            <Link to='category' params={{category: 16}}>
              Action
            </Link>
          </li>
          <li>
            <Link to='category' params={{category: 17}}>
              Adventure
            </Link>
          </li>
          <li>
            <Link to='category' params={{category: 18}}>
              Animation
            </Link>
          </li>
        </ul>
      </div>
    )
  },

});

module.exports = Categories;
