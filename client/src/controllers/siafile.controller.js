'use strict';

app.controller('SiafileCtrl', ['$scope', '$routeParams', 'SunfishSrvc', function($scope, $routeParams, SunfishSrvc){
    var params = $routeParams;
    SunfishSrvc.getSiafile(params.siafileId)
        .success(function(siafile){
            $scope.siafile = siafile;
            console.log($scope.siafile);
        });
}]);
