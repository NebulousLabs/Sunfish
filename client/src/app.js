'use strict';

var app = angular.module('sunfishApp', ['ngRoute']);

app.config(['$routeProvider', function($routeProvider) {
    $routeProvider
      .when('/', {
         templateUrl: 'views/main/main.html'
      })
      .when('/browse/', {
          controller: "BrowseCtrl",
          templateUrl: 'views/browse/browse.html'
      })
      .otherwise({redirectTo: '/'});
  }]);
