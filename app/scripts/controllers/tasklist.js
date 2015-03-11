'use strict';

/**
 * @ngdoc function
 * @name angDaysApp.controller:TasklistCtrl
 * @description
 * # TasklistCtrl
 * Controller of the angDaysApp
 */
angular.module('angDaysApp')
    .controller('TasklistCtrl', function ($scope, $http,shareTasks) {
        $http.get('/tasks')
            .success(function(data) {
                $scope.tasks = data ;
                shareTasks.settasks($scope.tasks);
            });
        $scope.hideContent = true;
        $scope.revealContent = function () {
            this.hideContent = !this.hideContent;
        };

        $scope.deleteTask = function(task) {
            $http.delete('/tasks', {params: {id:task.id}});
            
            $scope.tasks = $scope.tasks.filter(function(i) { return i.id !== task.id;});
        };
  });
