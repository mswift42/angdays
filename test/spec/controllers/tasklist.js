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
  });
});
