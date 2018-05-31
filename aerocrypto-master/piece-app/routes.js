//SPDX-License-Identifier: Apache-2.0

var controller = require('./controller.js');

module.exports = function(app){

  app.get('/get_piece/:id', function(req, res){
    controller.get_piece(req, res);
  });
  app.get('/add_piece/:piece', function(req, res){
    controller.add_piece(req, res);
  });
  app.get('/get_all_piece', function(req, res){
    controller.get_all_piece(req, res);
  });
}

