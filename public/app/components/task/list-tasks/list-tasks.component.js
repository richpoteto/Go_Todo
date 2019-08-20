angular.module('listTasks').component('listTasks', {
   templateUrl: 'components/task/list-tasks/list-tasks.tmpl.html',
    scope: {},
    controller: function($scope, toaster, taskResource) {
        let loadData = function() {
            $scope.tasks = taskResource.query();
            $scope.tasks.$promise.then(function () {
                toaster.pop('info', 'Success', 'Tasks loaded');
            }, function (response) {
                toaster.pop('error', response.data.title, response.data.description);
                console.log(JSON.stringify(response));
            });
        };
        $scope.deleteTask = function(id) {
            taskResource.delete({taskId: id}, function () {
                toaster.pop('info', 'Success', 'Task deleted');
                loadData();
            }, function (response) {
                toaster.pop('error', response.data.title, response.data.description);
                console.log(JSON.stringify(response));
            });
        };
        loadData();
    }
});
