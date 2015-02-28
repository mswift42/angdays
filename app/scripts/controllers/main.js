'use strict';

/**
 * @ngdoc function
 * @name angDaysApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the angDaysApp
 */
angular.module('angDaysApp')
  .controller('MainCtrl', function ($scope) {
    $scope.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
  });
