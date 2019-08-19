angular.module('listLabels').component('listLabels', {
   templateUrl: 'components/label/list-labels/list-labels.tmpl.html',
    scope: {},
    controller: function($scope, toaster, labelResource) {
        var loadData = function() {
            $scope.labels = labelResource.query();
            $scope.labels.$promise.then(function () {
                toaster.pop('info', 'Success', 'Labels loaded');
            }, function (response) {
                toaster.pop('error', response.statusText, response.description);
                console.log(JSON.stringify(response));
            });
        };
        $scope.deleteLabel = function(id) {
            labelResource.delete({labelId: id}, function () {
                toaster.pop('info', 'Success', 'Label deleted');
            }, function (response) {
                toaster.pop('error', response.data.title, response.data.description);
                console.log(JSON.stringify(response));
            });
            loadData();
        };
        loadData();
    }
});
