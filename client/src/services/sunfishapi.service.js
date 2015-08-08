'use strict';

app.factory('SunfishSrvc', ['$http', function($http){
    var baseUrl = '/api/siafile/';
    var sunfishSrvc = {};

    sunfishSrvc.getSiafiles = function(safeSearch) {
        return $http.get(baseUrl, {params: {'safe': safeSearch}});
    }

    sunfishSrvc.getSiafile = function(id) {
        return $http.get(baseUrl + id);
    }

    sunfishSrvc.upload = function(siafile) {
        return $http.post(baseUrl, siafile);
    }

    sunfishSrvc.searchSiafiles = function(searchString, safeSearch) {
        return $http.get(baseUrl + 'search', {
            params: {
                'tags': searchString,
                'safe': safeSearch
            }
        })
    }

    return sunfishSrvc;
}]);
