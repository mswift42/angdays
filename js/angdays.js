var daysApp = angular.module('daysApp', ['ngRoute', 'mgcrea.ngStrap', 'ngResource']);

daysApp.factory('Task', function($resource) {
    var Task = $resource('/tasks');
    return Task;
});
daysApp.factory('Edittask', function($resource) {
    return $resource('/tasks/:id', {},{
        update: {method: 'PUT'}
    });
});

daysApp.controller('ScrollCtrl',function($scope,$location,$anchorScroll) {
    $scope.scrollToTask = function(task) {
        $location.hash(task.id);
        $anchorScroll();
    };
});


daysApp.controller('TaskCtrl', function($scope, Task,Edittask,$filter,$resource) {
    Task.get(function(data) {
        $scope.tasks = data.tasks;
        $scope.agendas = data.agendaslice;
    });

    $scope.addTask = function(newtaskform) {
        if (newtaskform.$valid) {
            var task = new Task();
            task.summary = $scope.taskSummary;
            task.content = $scope.taskContent;
            task.scheduled = $scope.taskScheduled;
            task.done = 'Todo';
            task.$save();
            $scope.tasks.push(task);
            $scope.taskSummary = '';
            $scope.taskContent = '';
            $scope.taskScheduled = '';
            $scope.$apply(function() {
                $scope.showtask = true;
            });
//            $scope.$route.reload();
        }
    };


    $scope.change = function(task) {
        console.log(task);
        console.log(Edittask);
        console.log(typeof Edittask);
        console.log(Task);
        task.$save();



    };
    $scope.disabled = function(task) {
        return task.state !== undefined;
    };

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
        dateFormat: 'dd/MM/yyyy',
        startWeek: 1,
        autoClose: true
    });
});
