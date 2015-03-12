'use strict';

/**
 * @ngdoc function
 * @name angDaysApp.controller:NewtaskCtrl
 * @description
 * # NewtaskCtrl
 * Controller of the angDaysApp
 */
angular.module('angDaysApp')
    .controller('NewtaskCtrl', function ($scope,$http,shareTasks, dateService) {
        $scope.hideContent = true;
        $scope.revealContent = function () {
            $scope.hideContent = !$scope.hideContent;
        };
        $scope.saveTask = function() {
            
            var task = {"summary":$scope.formData.summary,
                        "content": $scope.formData.content};
            var sched = $scope.formData.scheduled;
            if (sched !== undefined) {
                task.scheduled = new Date(sched);
            } else {
                task.scheduled = dateService.nextWeek(new Date());
            }
            $http.post("/tasks",task).success(function() {
                                     $scope.hideContent = true;
                $scope.formData.summary = '';
                $scope.formData.content = '';
                $scope.formData.scheduled = '';
                shareTasks.add(task);
            });
            
        };
  });
