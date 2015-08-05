'use strict';

app.controller('UploadCtrl', ['$scope', '$location', 'SunfishSrvc', 'SiafileReaderSrvc', function($scope, $location, SunfishSrvc, SiafileReaderSrvc) {
    // Initialize the Siafile
    $scope.siafile = {};
    $scope.siafile.listed = true;
    $scope.siafile.safe = true;

    $scope.uploadSiafile = function() {
        // Read the client file for upload
        SiafileReaderSrvc.readfile().then(function(data) {
            // Encode the base64 file in ascii
            SiafileReaderSrvc.asciiEncode(data.base64, function(ascii){
                $scope.siafile.ascii = ascii;
                $scope.siafile.filename = data.filename;
                $scope.siafile.tags = $scope.tags.trim().split(",");

                // Process tags by making them stripped and lowercase
                for (var i = 0; i <$scope.siafile.tags.length; i++){
                    $scope.siafile.tags[i] = $scope.siafile.tags[i].trim().toLowerCase();
                };

                // Upload the processed siafile to the server
                SunfishSrvc.upload($scope.siafile)
                    .success(function(siafile) {
                        $location.path("/siafile/" + siafile.Id);
                    });
            });
        });
    };
}]);
