angular.module('listTasks').component('listTasks', {
   templateUrl: 'components/task/list-tasks/list-tasks.tmpl.html',
    scope: {},
    controller: function($scope, toaster, taskResource, labelResource) {
        $scope.labels = [];
        let loadTasks = function() {
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
                loadTasks();
            }, function (response) {
                toaster.pop('error', response.data.title, response.data.description);
                console.log(JSON.stringify(response));
            });
        };
        $scope.filterByLabels = function(task) {
            if ($scope.labels.length === 0) {
                return true;
            }
            if (task.labels === undefined || task.labels === null || task.labels.length === 0) {
                return false;
            }
            let inters = task.labels.filter(label => -1 !== $scope.labels.map(function (e) {
                return e.id;
            }).indexOf(label.id));
            return inters !== null && inters.length > 0;
        };
        loadTasks();
        $scope.order = {};
        $scope.selectedOrder = {};
        $scope.availableOrderBy = [
            {id: '-priority', label: 'Priority'},
            {id: 'due_date', label: 'Due date'},
        ];
        $scope.selectedOrderChanged = function () {
            $scope.order = $scope.selectedOrder.id;
        };
        $scope.clearSortFilter = function() {
            $scope.order = 'id';
        };
        $scope.orderBySettings = {
            selectionLimit: 1,
        };
    }
});
