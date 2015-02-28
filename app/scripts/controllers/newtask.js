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
        $scope.revealContent = function () {
            console.log($scope.showContent);
            $scope.showContent = !$scope.showContent;
        };

        
  });
