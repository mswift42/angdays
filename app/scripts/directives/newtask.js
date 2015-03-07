'use strict';

/**
 * @ngdoc directive
 * @name angDaysApp.directive:newTask
 * @description
 * # newTask
 */
angular.module('angDaysApp')
  .directive('newTask', function () {
    return {
        templateUrl: '../../views/newtask.html',
    };
  });
