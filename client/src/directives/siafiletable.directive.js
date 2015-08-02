'use strict';

app.directive('siafileTable', function() {
    var linkFunc = function(scope, element, attrs) {
            scope.sortOrder = '';

            scope.sortBy = function(columnName) {
                if (scope.sortOrder == columnName) {
                    scope.sortOrder = '-' + scope.sortOrder;
                } else {
                    scope.sortOrder = columnName;
                }
            };
        };
    return {
        templateUrl: 'views/siafiletable.html',
        link: linkFunc
    }
});
