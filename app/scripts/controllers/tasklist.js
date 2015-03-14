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
        $http.get('/api/tasks')
            .success(function(data) {
                $scope.tasks = data ;
                shareTasks.settasks($scope.tasks);
            });
        $scope.hideContent = true;
        $scope.revealContent = function () {
            this.hideContent = !this.hideContent;
        };

        $scope.deleteTask = function(task) {
            $http.delete('/api/tasks/' + task.id);
            $scope.tasks = $scope.tasks.filter(function(i)
                                               { return i.id !== task.id;});
        };
        $scope.editTask = function(task) {
            $http.post('/api/tasks/' + task.id, {data: {summary:task.summary,
                                                     content:task.content,
                                                     scheduled:new Date(task.scheduled),
                                                     done:task.done}});
                                  
            $scope.tasks = function() {
                for (var i = 0;i<$scope.tasks.length;i++) {
                    if ($scope.tasks[i].id === task.id) {
                        $scope.tasks[i] = task;
                    }
                }
            };
        };
        
    });
