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
        $scope.saveTask = function() {
            
            var task = {"summary":$scope.formData.summary}; 
            var sched = $scope.formData.scheduled;
            task.scheduled = (sched !== undefined || sched !== null) ? new Date(sched) : dateService.nextWeek(new Date());
            $http.post("/api/tasks",task).success(function() {
                                     $scope.hideContent = true;
                $scope.formData.summary = '';
                // $scope.formData.content = '';
                $scope.formData.scheduled = '';
                shareTasks.add(task);
            });
            
        };
  });
