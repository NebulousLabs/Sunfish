'use strict';

app.controller('SearchCtrl', ['$scope', 'SunfishSrvc', function($scope, SunfishSrvc){
    $scope.safe = true;

    $scope.search = function() {
        SunfishSrvc.searchSiafiles($scope.query, $scope.safe)
          .success(function(siafiles){
              $scope.siafiles = siafiles;
          });
    }

    // Function for filter to call on change
    $scope.updateSiafiles = function(){
        $scope.search();
    }
}]);
