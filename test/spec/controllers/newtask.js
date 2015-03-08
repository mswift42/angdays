'use strict';

describe('Controller: NewtaskCtrl', function () {

  // load the controller's module
  beforeEach(module('angDaysApp'));

  var NewtaskCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    NewtaskCtrl = $controller('NewtaskCtrl', {
      $scope: scope
    });
  }));

  it('should set hideContent initially to true.', function () {
      expect(scope.hideContent).toBeDefined();
      expect(scope.hideContent).toEqual(true);
  });
    it('revealContent should set hideContent to be false.', function() {
        expect(scope.revealContent).toBeDefined();
        scope.revealContent();
        expect(scope.hideContent).toEqual(false);
    });
    it('saveTask function should be defined', function() {
        expect(scope.saveTask).toBeDefined();
    });
});
