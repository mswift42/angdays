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
      template: '<div></div>',
      restrict: 'E',
      link: function postLink(scope, element, attrs) {
        element.text('this is the dblClickDirective directive');
      }
    };
  });
