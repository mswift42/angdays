'use strict';

/**
 * @ngdoc service
 * @name angDaysApp.dateService
 * @description
 * # dateService
 * Factory in the angDaysApp.
 */
angular.module('angDaysApp')
  .factory('dateService', function () {
    // Service logic
    // ...

    var meaningOfLife = 42;

    // Public API here
    return {
      someMethod: function () {
        return meaningOfLife;
      }
    };
  });
