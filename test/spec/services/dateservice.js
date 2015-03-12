'use strict';

describe('Service: dateService', function () {

  // load the service's module
  beforeEach(module('angDaysApp'));

  // instantiate service
  var dateService;
  beforeEach(inject(function (_dateService_) {
      dateService = _dateService_;
  }));

  it('should do something', function () {
      expect(!!dateService).toBe(true);
  });
    it('should return the correct Date for 31/03/2015', function() {
        var someday = new Date('March 10, 2015');
        var nextday = new Date('March 11, 2015');
        expect(dateService.nextWeek(someday)).toBe('17/03/2015');
        expect(dateService.nextWeek(nextday)).toBe('18/03/2015');
    });

});
