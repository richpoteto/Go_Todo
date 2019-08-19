angular.module('editTask').component('editTask', {
    templateUrl: 'components/task/edit-task/edit-task.tmpl.html',
    scope: {},
    controller: function($scope, $location, $routeParams, toaster, utils, taskResource) {
        $scope.editing = utils.isAnExistingOne();
        $scope.headerText = $scope.editing ? "Edit task" : "Create task";
        $scope.task = {labels: []};
        if ($scope.editing) {
            $scope.save = function () {
                console.log('$scope.task', JSON.stringify($scope.task));
                taskResource.update({taskId: $routeParams.taskId}, $scope.task).$promise.then(function () {
                    toaster.pop('info', 'Success', 'Task saved');
                    $location.path('/tasks');
                }, function (resp) {
                    toaster.pop('error', resp.data.title, resp.data.description);
                    console.error('response', JSON.stringify(resp));
                    $location.path('/tasks');
                });
            };
            taskResource.get({taskId: $routeParams.taskId}).$promise.then(function (task) {
                for (let property in task) {
                    if (task.hasOwnProperty(property)) {
                        if (property === 'labels') {
                            if (task[property] !== null) {
                                task[property].forEach(function (label) {
                                    $scope.task.labels.push(label);
                                }, this);
                            }
                            continue;
                        }
                        $scope.task[property] = task[property];
                    }
                }
                toaster.pop('info', 'Success', 'Task loaded');
            }, function (response) {
                toaster.pop('error', response.data.title, response.data.description);
                console.error('response', JSON.stringify(response));
            });
        } else {
            $scope.save = function () {
                taskResource.save($scope.task, function () {
                    toaster.pop('info', 'Success', 'Label saved');
                    $location.path('/tasks');
                }, function (resp) {
                    toaster.pop('error', resp.data.title, resp.data.description);
                    console.error('response', JSON.stringify(resp));
                    $location.path('/tasks');
                });
            };
        }
    }
});
