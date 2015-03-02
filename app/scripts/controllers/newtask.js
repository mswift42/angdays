'use strict';

/**
 * @ngdoc function
 * @name angDaysApp.controller:NewtaskCtrl
 * @description
 * # NewtaskCtrl
 * Controller of the angDaysApp
 */
angular.module('angDaysApp')
    .controller('NewtaskCtrl', function ($scope,$http) {
        $scope.hideContent = true;
        $scope.revealContent = function () {
            $scope.hideContent = !$scope.hideContent;
        };
        $scope.datepickeroptions = {
            format: 'yyyy-mm-dd'
        };
        $scope.saveTask = function() {
            console.log($scope.newtaskform);
            $http.post("/tasks",{"summary":$scope.formData.summary,
                                 "content":$scope.formData.content,
                                 "scheduled":$scope.formData.scheduled}).success(function() {
                                     $scope.hideContent = true;
                                     $scope.formData.summary = '';});
            
        };
  });
