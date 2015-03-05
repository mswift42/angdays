'use strict';

describe('Directive: dblClickDirective', function () {

  // load the directive's module
  beforeEach(module('angDaysApp'));

  var element,
    scope;

  beforeEach(inject(function ($rootScope) {
    scope = $rootScope.$new();
  }));

  it('should make hidden element visible', inject(function ($compile) {
    element = angular.element('<dbl-click-directive></dbl-click-directive>');
    element = $compile(element)(scope);
    expect(element.text()).toBe('this is the dblClickDirective directive');
  }));
});
