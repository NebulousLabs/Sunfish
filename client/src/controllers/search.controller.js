'use strict';

app.controller('SearchCtrl', ['$scope', 'SunfishSrvc', function($scope, SunfishSrvc){
    $scope.search = function() {
        SunfishSrvc.searchSiafiles($scope.query)
          .success(function(siafiles){
              $scope.siafiles = siafiles;
          });
    }
}]);
