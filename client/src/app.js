'use strict';

var app = angular.module('sunfishApp', ['ngRoute']);
'$compileProvider',
    app.config(['$routeProvider', '$compileProvider', function($routeProvider, $compileProvider) {
        $routeProvider
        .when('/', {
            templateUrl: 'views/main/main.html',
        })
        .when('/browse/', {
            controller: "BrowseCtrl",
            templateUrl: 'views/browse/browse.html',
        })
        .when('/search/', {
            controller: "SearchCtrl",
            templateUrl: 'views/search/search.html',
        })
        .when('/upload/', {
            controller: "UploadCtrl",
            templateUrl: 'views/upload/upload.html',
        })
        .otherwise({redirectTo: '/'});

        // Whitelist allowed URL types
        $compileProvider.aHrefSanitizationWhitelist(/^\s*(https?|ftp|mailto|data):/);
    },
    ]);
