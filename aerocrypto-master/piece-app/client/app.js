// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();
	
	$scope.queryAllPiece = function(){

		appFactory.queryAllPiece(function(data){
			console.log(data);
			var array = [];
			for (var i = 0; i < data.length; i++){
				parseInt(data[i].Key);
				data[i].Record.Key = parseInt(data[i].Key);
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return parseFloat(a.Key) - parseFloat(b.Key);
			});

			console.log(array);
			$scope.all_piece = array;
		});
	}

	$scope.queryPiece = function(){

		var id = $scope.piece_id;

		appFactory.queryPiece(id, function(data){
			$scope.query_piece = data;

			if ($scope.query_piece == "Could not locate piece"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	$scope.recordPiece = function(){

		appFactory.recordPiece($scope.piece, function(data){
			$scope.create_piece = data;
			$("#success_create").show();
		});
	}

});

// Angular Factory
app.factory('appFactory', function($http){
	
	var factory = {};

    factory.queryAllPiece = function(callback){

    	$http.get('/get_all_piece/').success(function(output){
			callback(output)
		});
	}

	factory.queryPiece = function(id, callback){
    	$http.get('/get_piece/'+id).success(function(output){
			callback(output)
		});
	}

	factory.recordPiece = function(data, callback){

		var piece = data.id + "-" + data.refPiece + "-" + data.nomPiece 
		+ "-" + data.nomFabricant + "-" + data.numeroLot + "-" + data.dateExp 
		+ "-" + data.localisationVille + "-" + data.localisationEtablissement 
		+ "-" + data.dateCreation + "-" + data.dateReception + "-" + "false";



   	$http.get('/add_piece/'+ piece).success(function(output){
			callback(output)
		});
	}


	return factory;
});



