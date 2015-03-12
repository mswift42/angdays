describe('Given I am adding a new task', function() {
    describe('Focussing the text input should reveal the complete task form.', function() {
        beforeEach(function() {
            browser.get('/#/');

            var summaryInput = $('input');
            summaryInput.sendKeys('first task');

        });
        it('should reveal task form', function() {
            var tcont = $('.tcontent');
            expect(tcont.getCssValue).toNotEqual('none');
        });
    });
});
