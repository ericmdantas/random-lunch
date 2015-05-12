"use strict";

describe('place model', function() {
    var _Place;
    var _placeComplete;

    beforeEach(module('randomLunch'));

    beforeEach(inject(function($injector)
    {
        _Place = $injector.get('Place');

        _placeComplete =
        {
            name: 'ABCDEF',
            address: '231213',
            image: 'abhjsdjsdhsd16456'
        }
    }));

    describe('init', function()
    {
        it('should have right props for instance', function()
        {
            var _p = new _Place();

            expect(_p.name).toBeNull();
            expect(_p.address).toBeNull();
            expect(_p.image).toBeNull();
        })

        it('should have the right info when instantiated - constructor', function()
        {
            var _p = new _Place(_placeComplete);

            expect(angular.equals(_p, _placeComplete)).toBeTruthy();
        })
    })

    describe('isValid', function() {
        it('should return false - no name', function() {
          var _p = new _Place(_placeComplete);

          delete _p.name;

          expect(_p.isValid()).toBeFalsy();
        })

        it('should return false - no address', function()
        {
            var _p = new _Place(_placeComplete);

            delete _p.address;

            expect(_p.isValid()).toBeFalsy();
        })

        it('should return true - no image', function() {
          var _p = new _Place(_placeComplete);

          delete _p.image;

          expect(_p.isValid()).toBeTruthy();

        })

        it('should return true - complete', function()
        {
            var _p = new _Place(_placeComplete);

            expect(_p.isValid()).toBeTruthy();
        })
    })
})
