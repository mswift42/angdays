'use strict';

/**
 * @ngdoc directive
 * @name angDaysApp.directive:taskList
 * @description
 * # taskList
 */
angular.module('angDaysApp')
  .directive('taskList', function () {
    return {
      templateUrl: '../../views/tasklist.html'
    };
  });
