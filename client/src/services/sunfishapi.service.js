'use strict';

app.factory('SunfishSrvc', ['$http', function($http){
    var baseUrl = '/api/siafile/';
    var sunfishSrvc = {};

    sunfishSrvc.getSiafiles = function() {
        return $http.get(baseUrl);
    }

    sunfishSrvc.upload = function(siafile) {
        return $http.post(baseUrl, siafile);
    }

    return sunfishSrvc;
}]);
