;(function(angular)
{
  "use strict";

 angular
    .module('randomLunch')
    .factory('Place', [function()
    {
      var Place = function(place) {
          this.name = null;
          this.address = null;
          this.image = null;

          angular.extend(this, place);
      }

      Place.prototype.isValid = function() {
          return !!this.name && this.address;
      }

      return Place;

    }]);

}(window.angular));
