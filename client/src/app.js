'use strict';

var app = angular.module('sunfishApp', ['ngRoute']);
'$compileProvider',
    app.config(['$routeProvider', '$compileProvider', function($routeProvider, $compileProvider) {
        $routeProvider
        .when('/', {
            templateUrl: 'views/main.html',
        })
        .when('/browse/', {
            controller: "BrowseCtrl",
            templateUrl: 'views/browse.html',
        })
        .when('/search/', {
            controller: "SearchCtrl",
            templateUrl: 'views/search.html',
        })
        .when('/upload/', {
            controller: "UploadCtrl",
            templateUrl: 'views/upload.html',
        })
        .when('/siafile/:siafileId', {
            controller: "SiafileCtrl",
            templateUrl: 'views/siafile.html',
        })
        .otherwise({redirectTo: '/'});

        // Whitelist allowed URL types
        $compileProvider.aHrefSanitizationWhitelist(/^\s*(https?|ftp|mailto|data):/);
    }]);

// Initialize foundation once the page has loaded.
app.run(function($timeout){
  $timeout(function() {
    $(document).foundation();
  }, 500);
});
