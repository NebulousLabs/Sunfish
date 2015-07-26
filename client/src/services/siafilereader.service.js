'use strict';

app.factory('SiafileReaderSrvc', ['$q', function($q){
    var readerSrvc = {};

    readerSrvc.asciiEncode = function(base64, callback){
        var ascii = base64.replace(/\+/g, '-').replace(/\//g, '_');
        callback(ascii);
    }

    readerSrvc.urlEncode = function(ascii, callback){
        var dataUrl = ascii.replace(/-/g, '+').replace(/_/g, '/');
        callback(dataUrl);
    }

    readerSrvc.readfile = function(){
        var deferred = $q.defer();

        var file = document.querySelector('input[type=file]').files[0];
        var reader = new FileReader();

        reader.onloadend = function () {
            var filename = file.name;
            if(filename.indexOf(".sia") != -1){
                deferred.resolve({'base64': btoa(reader.result), 'filename': filename});
            } else {
                alert("Error: Not a Siafile!");
            }
        }

        if (file) {
            reader.readAsBinaryString(file);
        }
        return deferred.promise;
    }

    return readerSrvc;
}]);
