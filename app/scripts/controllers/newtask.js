'use strict';

/**
 * @ngdoc function
 * @name angDaysApp.controller:NewtaskCtrl
 * @description
 * # NewtaskCtrl
 * Controller of the angDaysApp
 */
angular.module('angDaysApp')
    .controller('NewtaskCtrl', function ($scope) {
        $scope.hideContent = true;
        $scope.revealContent = function () {
            $scope.hideContent = !$scope.hideContent;
        };
  });
