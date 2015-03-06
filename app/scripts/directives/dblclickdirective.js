'use strict';

/**
 * @ngdoc directive
 * @name angDaysApp.directive:dblClickDirective
 * @description
 * # dblClickDirective
 */
angular.module('angDaysApp')
    .directive('dblClickDirective', function () {
        return {
            scope: {},
            restrict: 'A',
            link: function postLink(scope, element, attrs) {
                element.text('this is the dblClickDirective directive');
            }
        };
    });
