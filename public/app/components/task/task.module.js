angular.module('taskModule', ['listTasks', 'editTask']).
factory('taskResource', ['$resource',
    function($resource) {
        return $resource('/api/v1.0/tasks/:taskId', {labelId: '@taskId'}, {
            update: {
                method: 'PUT'
            }
        });
    }
]);
