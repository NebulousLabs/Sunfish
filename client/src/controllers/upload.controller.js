'use strict';

app.controller('UploadCtrl', ['$scope', 'SunfishSrvc', 'SiafileReaderSrvc', function($scope, SunfishSrvc, SiafileReaderSrvc) {
    $scope.siafile = {};

    $scope.uploadSiafile = function() {
        SiafileReaderSrvc.readfile().then(function(data) {
            SiafileReaderSrvc.asciiEncode(data.base64, function(ascii){
                $scope.siafile.ascii = ascii;
                $scope.siafile.filename = data.filename;

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
