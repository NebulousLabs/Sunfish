'use strict';

app.controller('siafileRowCtrl', function($scope){
  // Used to find number of tags to show based on their length and a given max
  $scope.tagsCount = function(tags, maxLength){
      var count = 0;
      var totalLength = 0;
      for (var i = 0; i < tags.length; i++){
          if (totalLength + tags[i].length < maxLength) {
              count += 1;
              totalLength += tags[i].length;
          } else {
              return count;
          }
      }
      return count;
  }
});
