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

      var datefactory = {};
      
      datefactory.formatDate = function(date) {
          return date.toLocaleString().split(' ')[0];
      };

      datefactory.nextWeek = function(day) {
          var dayms = day.getTime();
          return datefactory.formatDate(dayms + (7 * 24 * 60 * 60 * 1000));
      };

      return datefactory;
          
          

    // Public API here
  });
