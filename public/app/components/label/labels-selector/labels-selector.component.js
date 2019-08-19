angular.module('labelsSelector', ['angularjs-dropdown-multiselect']).component('labelsSelector', {
    bindings: {
        selectedLabels: '='
    },
    templateUrl: 'components/label/labels-selector/labels-selector.tmpl.html',
    controller: function ($scope, toaster, labelResource) {
        let self = this;
        this.$onInit = function() {
            $scope.selectedLabels = self.selectedLabels;
        };
        $scope.availableLabels = [];
        $scope.labels = labelResource.query();
        $scope.labels.$promise.then(function (response) {
            response.forEach(function (label) {
                $scope.availableLabels.push(label);
            }, this);
        }, function (response) {
            console.log(JSON.stringify(response));
        });
        $scope.dropdownSettings = {
            checkBoxes: true,
            styleActive: true,
            searchField: 'name',
            enableSearch: true,
            buttonDefaultText: 'Select labels',
            idProperty: "id",
            selectedToTop: true,
            scrollableHeight: '250px',
            scrollable: true,
            smartButtonMaxItems: 4,
            displayProp: 'name'
        };
    }
});
