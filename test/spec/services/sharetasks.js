'use strict';

describe('Service: shareTasks', function () {

  // load the service's module
  beforeEach(module('angDaysApp'));

  // instantiate service
  var shareTasks;
  beforeEach(inject(function (_shareTasks_) {
    shareTasks = _shareTasks_;
  }));

  it('should do something', function () {
    expect(!!shareTasks).toBe(true);
  });

});
