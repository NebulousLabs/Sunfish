'use strict';

angular.module('sunfishApp')
  .directive('siafileRow', function() {
    return {
      restrict: 'E',
      scope: {
        title: '@',
        description: '@',
        filename: '@',
        fileData: '@',
        tags: '@',
      }
    }
  });
