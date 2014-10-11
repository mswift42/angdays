var daysApp = angular.module('daysApp', ['ngRoute', 'mgcrea.ngStrap', 'ngResource']);

daysApp.factory('Task', function($resource) {
    'use strict';
    var Task = $resource('/tasks');
    return Task;
});
daysApp.factory('Edittask', function($resource) {
    'use strict';
    return $resource('/tasks/:id', {},{
        update: {method: 'POST'}
    });
});

daysApp.controller('ScrollCtrl',function($scope,$location,$anchorScroll) {
    'use strict';
    $scope.scrollToTask = function(task) {
        $location.hash(task.id);
        $anchorScroll();
    };
});


daysApp.controller('TaskCtrl', function(Task,Edittask, $scope)  {
    'use strict';
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
            $scope.tasks.push(task.tasks);
            $scope.taskSummary = '';
            $scope.taskContent = '';
            $scope.taskScheduled = '';
            // $scope.$apply(function() {
            //     $scope.showtask = true;
            // });
//            $scope.$route.reload();
        }
    };
    $scope.update = function(task) {
        $scope.task.summary = task.summary;
        $scope.task.content = task.content;
        $scope.task.scheduled = task.scheduled;
        $scope.task.done = task.done;
        Edittask.update({id:task.id,summary:task.summary,content:task.content,
                         done:task.done,scheduled:task.scheduled});
        $scope.$apply(function() {
            $scope.showedit = false;
        });
    };

    $scope.disabled = function(task) {
        return task.state !== undefined;
    };

});



daysApp.config(function($routeProvider) {
    'use strict';
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
daysApp.directive('navbarHeader', function() {
    'use strict';
    return {
        restrict: 'EA',
        templateUrl: 'partials/navbar.html'
    };
});
daysApp.config(function($datepickerProvider) {
    'use strict';
    angular.extend($datepickerProvider.defaults, {
        dateFormat: 'dd/MM/yyyy',
        startWeek: 1,
        autoClose: true
    });
});
