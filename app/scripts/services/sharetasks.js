'use strict';

/**
 * @ngdoc service
 * @name angDaysApp.shareTasks
 * @description
 * # shareTasks
 * Factory in the angDaysApp.
 */
angular.module('angDaysApp')
  .factory('shareTasks', function () {
    // Service logic
    // ...
      var tasks = [];
      var taskService = {};
      taskService.add = function(item) {
          tasks.unshift(item);
      };
      taskService.list = function() {
          return tasks;
      };
      taskService.settasks = function(items) {
          tasks = items;
      };

    // Public API here
      return taskService; 
  });
