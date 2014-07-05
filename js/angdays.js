var daysApp = angular.module('daysApp', ['ngRoute', 'mgcrea.ngStrap', 'ngResource']);

daysApp.factory('Task', function($resource) {
    var Task = $resource('/tasks');
    return Task;
});
daysApp.controller('TaskCtrl', function($scope, Task) {
    Task.get(function(data) {
        $scope.tasks = data.tasks;
        $scope.agenda = data.agendaslice;
    });

    $scope.addTask = function(newtaskform) {
        if (newtaskform.$valid) {
            var task = new Task();
            task.summary = $scope.taskSummary;
            task.content = $scope.taskContent;
            task.scheduled = $scope.taskScheduled;
            task.done = "Todo";
            task.$save();
            task.state =  'saving';
            $scope.tasks.push(task);
            $scope.taskSummary = '';
            $scope.taskContent = '';
            $scope.taskScheduled = '';
            $scope.$apply(function() {
                $scope.showtask = true;
            });
            $scope.$route.reload();
            task.updating = true;
            task.state = 'updating';

        }
    };
    $scope.change = function(task) {
        task.$save();
        task.state = 'updating';
    };
    $scope.disabled = function(task) {
        return task.state !== undefined;
    };
    // $scope.archive = function() {
    //     Task.remove(function() {
    //         Task.query(function(tasks) {
    //             $scope.tasks = tasks;
    //         });
    //     });
    // };
});



daysApp.config(function($routeProvider) {
    $routeProvider.
        when('/', {
            templateUrl: 'partials/tasks.html',
            controller: 'TaskCtrl'
        }).
        when('/about', {
            templateUrl: 'partials/about.html'
        }).
        otherwise({
            redirectTo: '/'
        });
});
daysApp.directive("navbarHeader",function() {
    return {
        restrict: 'EA',
        templateUrl: "partials/navbar.html"
    };
});
daysApp.config(function($datepickerProvider) {
    angular.extend($datepickerProvider.defaults, {
        dateFormat: 'dd/mm/yyyy',
        startWeek: 1
    });
});
