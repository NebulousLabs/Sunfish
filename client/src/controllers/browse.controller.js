'use strict';

app.controller('BrowseCtrl', ['$scope', 'SunfishSrvc', function($scope, SunfishSrvc) {
    $scope.siafiles = [];

    SunfishSrvc.getSiafiles()
        .success(function(siafiles){
            $scope.siafiles = siafiles;
        });
}]);
