'use strict';

app.controller('UploadCtrl', ['$scope', 'SunfishSrvc', 'SiafileReaderSrvc', function($scope, SunfishSrvc, SiafileReaderSrvc) {
    $scope.siafile = {};

    $scope.uploadSiafile = function() {
        SiafileReaderSrvc.readfile().then(function(data) {
            SiafileReaderSrvc.asciiEncode(data.base64, function(ascii){
                $scope.siafile.ascii = ascii;
                $scope.siafile.filename = data.filename;
                $scope.siafile.tags = $scope.tags.trim().split(",");

                for (var i = 0; i <$scope.siafile.tags.length; i++){
                    $scope.siafile.tags[i] = $scope.siafile.tags[i].trim().toLowerCase();
                };

                SunfishSrvc.upload($scope.siafile)
                .success(function(siafile) {
                    console.log(siafile);
                })
                .error(function(error) {
                    console.log(error);
                });
            });
        });
    };
}]);
