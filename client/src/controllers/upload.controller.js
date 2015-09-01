'use strict';

app.controller('UploadCtrl', ['$scope', '$location', 'SunfishSrvc', 'SiafileReaderSrvc', function($scope, $location, SunfishSrvc, SiafileReaderSrvc) {
    // Initialize the Siafile
    $scope.siafile = {};
    $scope.siafile.listed = true;
    $scope.siafile.safe = true;
    $scope.safeClicked = false;
    $scope.asciiUpload = true;

    $scope.tags = [];
    $scope.$watch('tags.length', function() {
        var safe = true;
        for (var i = 0; i < $scope.tags.length; i++) {
            if ($scope.tags[i].text.toLowerCase() === 'nsfw'){
                safe = false;
            }
        }

        if (safe === true && !$scope.safeClicked) {
            $scope.siafile.safe = true;
        } else {
            $scope.siafile.safe = false;
        }
    });

    $scope.uploadSiafile = function() {
        var handleAscii = function(ascii, filename) {
            $scope.siafile.ascii = ascii;
            $scope.siafile.filename = filename;
            $scope.siafile.tags = [];

            // Process tags by making them stripped and lowercase
            for (var i = 0; i <$scope.tags.length; i++){
                $scope.siafile.tags.push($scope.tags[i].text.trim().toLowerCase());
            };

            // Upload the processed siafile to the server
            SunfishSrvc.upload($scope.siafile)
                .success(function(siafile) {
                    $location.path("/siafile/" + siafile.Id);
                });
        }

        // If user uploaded with ascii paste
        if ($scope.asciiUpload) {
            handleAscii($scope.siafileAscii, $scope.siafileName);
        } else {
            // Read the client file for upload
            SiafileReaderSrvc.readfile().then(function(data) {
                // Encode the base64 file in ascii
                SiafileReaderSrvc.asciiEncode(data.base64, function(ascii){
                    handleAscii(ascii, data.filename);
                });
            });
        }
    };

    $scope.setUploadType = function(bool) {
        $scope.asciiUpload = bool;
    }
}]);
