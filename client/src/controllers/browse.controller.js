'use strict';

app.controller('BrowseCtrl', ['$scope', 'SunfishSrvc', function($scope, SunfishSrvc) {
    $scope.siafiles = [];
    $scope.safe = true;

    var getSiafiles = function(){
        SunfishSrvc.getSiafiles($scope.safe).success(function(siafiles){
            $scope.siafiles = siafiles;
        });
    }

    getSiafiles();

    $scope.updateSiafiles = function (){
        getSiafiles();
    }
}]);
