var express = require('express');
var router = express.Router();
const title = process.env.VAL || 'Express';

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title });
});

module.exports = router;
