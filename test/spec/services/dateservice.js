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
        expect(dateService.nextWeek(someday).getMonth()).toBe(2);
        expect(dateService.nextWeek(nextday).getDate()).toBe(18);
    });
    it('should parse dates of format dd/mm/yyyy',function() {
        var day1 = "31/03/2000";
        var day2 = "04/02/2000";
        expect(dateService.parseDate(day1).getMonth()).toBe(2);
        expect(dateService.parseDate(day2).getDate()).toBe(4);
    });

});
