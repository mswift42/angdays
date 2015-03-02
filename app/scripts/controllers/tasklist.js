'use strict';

/**
 * @ngdoc function
 * @name angDaysApp.controller:TasklistCtrl
 * @description
 * # TasklistCtrl
 * Controller of the angDaysApp
 */
angular.module('angDaysApp')
  .controller('TasklistCtrl', function ($scope) {
    $scope.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
  });
