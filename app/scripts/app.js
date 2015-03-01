'use strict';

/**
 * @ngdoc overview
 * @name angDaysApp
 * @description
 * # angDaysApp
 *
 * Main module of the application.
 */
angular
  .module('angDaysApp', [
    'ngAnimate',
    'ngCookies',
    'ngResource',
    'ngRoute',
    'ngSanitize',
      'ngTouch',
      'angular-datepicker'
  ])
  .config(function ($routeProvider) {
    $routeProvider
      .when('/', {
        templateUrl: 'views/main.html',
        controller: 'MainCtrl'
      })
      .when('/about', {
        templateUrl: 'views/about.html',
        controller: 'AboutCtrl'
      })
      .otherwise({
        redirectTo: '/'
      });
  });
