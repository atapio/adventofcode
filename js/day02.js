var day02 = require("./lib/day02");

var myAnswer = function(callback) {
  var requiredPaper = 0;
  var requiredRibbon = 0;

  var rl = require('readline').createInterface({
    input: require('fs').createReadStream('day02.input')
  });

  rl.on('line', function(line) {
    requiredPaper += day02.wrappingPaper(line);
    requiredRibbon += day02.ribbon(line);
  });

  rl.on('close', function() {
    callback([requiredPaper, requiredRibbon]);
  });
}

myAnswer(console.log);
