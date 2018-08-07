var express = require('express');
const request = require('request');
var router = express.Router();


/* GET home page. */
router.get('/', function (req, res, next) {
    res.render('index');
});

router.get('/dashboard', function (req, res, next) {

    request('http://localhost:3000/test', function (err, body) {
        if (err) {
            res.send(err)
        } else {
            var viewData = JSON.parse(body.body)
            res.render('dashboard', viewData);
        }
    });

});

module.exports = router;