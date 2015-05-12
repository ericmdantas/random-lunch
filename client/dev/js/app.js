;(function(angular)
{
  "use strict";

  angular
    .module('randomLunch', [
      'ngResource',
      'ngNewRouter',
      'ngAnimate',
      'emd.ng-xtorage'])
    .config(['$locationProvider', function($locationProvider)
    {
      $locationProvider.html5Mode(true);
    }]);

}(window.angular));

