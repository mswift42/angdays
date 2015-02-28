'use strict';

/**
 * @ngdoc function
 * @name angDaysApp.controller:AboutCtrl
 * @description
 * # AboutCtrl
 * Controller of the angDaysApp
 */
angular.module('angDaysApp')
  .controller('AboutCtrl', function ($scope) {
    $scope.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
  });
