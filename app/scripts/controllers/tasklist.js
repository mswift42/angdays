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
            $http({
                url: '/api/tasks/' + task.id,
                method: 'POST',
                data: {'summary':task.summary,
                       'content':task.content,
                       'id':task.id,
                       'scheduled': new Date(task.scheduled),
                       'done':task.done}
            }).then(function(_) {
                shareTasks.settasks($scope.tasks);
            });
            this.hideContent = !this.hideContent;
        };
    });
