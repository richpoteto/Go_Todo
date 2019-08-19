'use strict';

angular.module('utilsModule', []).service('utils', ['$location', function ($location) {
    this.isAnExistingOne = function() {
        var exist = false;
        var path = $location.path();
        if (path !== null) {
            var pl = path.split('/');
            var last = pl[pl.length - 1];
            var num = parseInt(last, 10);
            exist = !isNaN(num);
        }
        return exist;
    };
}]);

// Declare app level module which depends on views, and core components
angular.module('myApp', [
  'ngRoute',
  'ngAnimate',
  'myApp.view1',
  'myApp.view2',
  'myApp.version',
  'labelModule',
  'taskModule',
  'utilsModule'
]).
config(['$locationProvider', '$routeProvider', function($locationProvider, $routeProvider) {
  $locationProvider.hashPrefix('!');

  $routeProvider
      .when('/labels', {
          template: '<list-labels></list-labels>'
      })
      .when('/label/:labelId', {
          template: '<edit-label></edit-label>',
      })
      .when('/label', {
          template: '<edit-label></edit-label>',
      })
      .when('/tasks', {
          template: '<list-tasks></list-tasks>'
      })
      .when('/task/:taskId', {
          template: '<edit-task></edit-task>',
      })
      .when('/task', {
          template: '<edit-task></edit-task>',
      })
      .otherwise({redirectTo: '/view1'});
}]);
