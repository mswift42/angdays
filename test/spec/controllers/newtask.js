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

  it('should attach a list of awesomeThings to the scope', function () {
    expect(scope.awesomeThings.length).toBe(3);
  });
});
