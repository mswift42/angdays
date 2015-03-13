'use strict';

describe('Controller: TasklistCtrl', function () {

  // load the controller's module
  beforeEach(module('angDaysApp'));

  var TasklistCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
      scope = $rootScope.$new();
      
    TasklistCtrl = $controller('TasklistCtrl', {
        $scope: scope
    });
  }));
    

    it('hideContent should be set to true', function () {
        expect(scope.hideContent).toBeDefined();
        expect(scope.hideContent).toBe(true);
        scope.revealContent();
        expect(scope.hideContent).toBe(false);
    });
    
    it('should delete a task from scope.tasks', function() {
     scope.tasks = [{"summary":"task1","id":"123"},
                   {"summary": "task2", "id":"124"}];
        scope.deleteTask(scope.tasks[0]);
        expect(scope.tasks.length).toBe(1);
        expect(scope.tasks[0].id).toBe("124");
    });
    it('should edit a task with a given id', function() {
        scope.tasks = [{"summary":"task1", "id":"123"},
                       {"summary":"task2", "id": "124"}];
        scope.editTask({"summary":"taskone","id":"123"});
        expect(scope.tasks[0].summary).toBe("taskone");
    });
});
