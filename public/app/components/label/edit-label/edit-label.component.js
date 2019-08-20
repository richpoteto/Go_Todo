angular.module('editLabel').component('editLabel', {
   templateUrl: 'components/label/edit-label/edit-label.tmpl.html',
    scope: {},
    controller: function($scope, $location, $routeParams, toaster, utils, labelResource) {
        $scope.editing = utils.isAnExistingOne();
        $scope.headerText = $scope.editing ? "Edit label" : "Create label";
        $scope.label = {};
        if ($scope.editing) {
            $scope.save = function () {
                labelResource.update({labelId: $routeParams.labelId}, $scope.label).$promise.then(function () {
                    toaster.pop('info', 'Success', 'Label saved');
                    $location.path('/labels');
                }, function (resp) {
                    toaster.pop('error', resp.data.title, resp.data.description);
                    console.error('response', JSON.stringify(resp));
                });
            };
            labelResource.get({labelId: $routeParams.labelId}).$promise.then(function (label) {
                for (let property in label) {
                    if (label.hasOwnProperty(property)) {
                        $scope.label[property] = label[property];
                    }
                }
                toaster.pop('info', 'Success', 'Label loaded');
            }, function (response) {
                toaster.pop('error', response.data.title, response.data.description);
                console.error('response', JSON.stringify(response));
            });
        } else {
            $scope.save = function () {
                labelResource.save($scope.label, function () {
                    toaster.pop('info', 'Success', 'Label saved');
                    $location.path('/labels');
                }, function (resp) {
                    toaster.pop('error', resp.data.title, resp.data.description);
                    console.error('response', JSON.stringify(resp));
                });
            };
        }
    }
});
