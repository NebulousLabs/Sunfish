'use strict';

app.controller('SiafileCtrl', ['$scope', '$routeParams', 'SunfishSrvc', 'SiafileReaderSrvc', function($scope, $routeParams, SunfishSrvc, SiafileReaderSrvc){
    var params = $routeParams;

    SunfishSrvc.getSiafile(params.siafileId)
        .success(function(siafile){
            $scope.siafile = siafile;
            SiafileReaderSrvc.urlEncode(siafile.ascii, function(urlString) {
                $scope.downloadUrl = 'data:application/sia;base64,' + urlString;
            });
        });
  $scope.copyASCII = function(){
      var asciiText = $( "input[name='asciiText'" );
      asciiText.show();
      asciiText.select();
  }
}]);
