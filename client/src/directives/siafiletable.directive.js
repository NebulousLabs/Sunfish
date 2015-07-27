'use strict';

app.directive('siafileTable', function() {
    return {
        templateUrl: 'views/siafiletable/siafiletable.html',
        link: function(scope, element, attrs) {
            scope.sortOrder = '';

            scope.sortBy = function(columnName) {
                if (scope.sortOrder == columnName) {
                    scope.sortOrder = '-' + scope.sortOrder;
                } else {
                    scope.sortOrder = columnName;
                }
            };
        },
    }
});

