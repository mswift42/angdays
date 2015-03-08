'use strict';

/**
 * @ngdoc function
 * @name angDaysApp.controller:NewtaskCtrl
 * @description
 * # NewtaskCtrl
 * Controller of the angDaysApp
 */
angular.module('angDaysApp')
    .controller('NewtaskCtrl', function ($scope,$http,shareTasks) {
        $scope.hideContent = true;
        $scope.revealContent = function () {
            $scope.hideContent = !$scope.hideContent;
        };
        $scope.saveTask = function() {
            var task = {"summary":$scope.formData.summary,
                        "content": $scope.formData.content,
                        "scheduled": $scope.formData.scheduled};
            $http.post("/tasks",task).success(function() {
                                     $scope.hideContent = true;
                $scope.formData.summary = '';
                shareTasks.add(task);
            });
            
        };
  });
