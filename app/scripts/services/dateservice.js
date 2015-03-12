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
      
      datefactory.parseDate = function(datestring) {
          var spl = datestring.split('/');
          return new Date(spl[2],spl[1],spl[0]);
      };

      datefactory.nextWeek = function(day) {
          var dayms = day.getTime();
          return new Date(dayms + (7 * 24 * 60 * 60 * 1000));
      };

      return datefactory;
          
          

    // Public API here
  });
