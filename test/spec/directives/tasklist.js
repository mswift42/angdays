'use strict';

describe('Directive: taskList', function () {

  // load the directive's module
  beforeEach(module('angDaysApp'));

  var element,
    scope;

  beforeEach(inject(function ($rootScope) {
    scope = $rootScope.$new();
  }));

  it('should make hidden element visible', inject(function ($compile) {
    element = angular.element('<task-list></task-list>');
    element = $compile(element)(scope);
    expect(element.text()).toBe('this is the taskList directive');
  }));
});
