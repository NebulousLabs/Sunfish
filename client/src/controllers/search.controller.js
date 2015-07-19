'use strict';

app.controller('SearchCtrl', ['$scope', 'SunfishSrvc', function($scope, SunfishSrvc){
    SunfishSrvc.getSiafiles()
        .success(function(siafiles) {
        });
}]);
