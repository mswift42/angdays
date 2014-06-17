var daysApp = angular.module('daysApp',['ngRoute','mgcrea.ngStrap']);

daysApp.factory('Task',function($resource) {
  var Task = $resource('/tasks');
  return Task;
});
daysApp.controller('TaskCtrl', function($scope, Task) {
  $scope.tasks = Task.query();
  $scope.addTask = function() {
    var task = new Task();
    task.summary = $scope.task.summary;
    task.content = $scope.task.content;
    task.scheduled = $scope.task.scheduled;
    task.$save();
    task.state =  'saving';
    $scope.tasks.push(task);
    $scope.task.summary = '';
    task.updating = true;
    task.done = false;
  };
  $scope.change = function(task) {
    task.$save();
    task.state = 'updating';
  };
  $scope.disabled = function(task) {
    return task.state != undefined;
  };
});

daysApp.controller('DaysController',function($scope) {
  var message = "hello";
});
daysApp.controller('NewTaskController',function($scope) {
  $scope.message = "newtask text";
});
daysApp.config(function($routeProvider) {
  $routeProvider.
    when('/newtask', {
      templateUrl: 'partials/newtask.html',
      controller: 'NewTaskController'
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
    startWeek: 1
  });
});
