angular.module('labelModule', ['listLabels', 'editLabel']).
factory('labelResource', ['$resource',
    function($resource) {
        return $resource('/api/v1.0/labels/:labelId', {labelId: '@labelId'}, {
            update: {
                method: 'PUT'
            }
        });
    }
]);
