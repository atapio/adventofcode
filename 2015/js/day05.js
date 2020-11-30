var day05 = require("./lib/day05");

var myAnswer = function(callback) {
  var niceCount = 0;
  var niceCountPartTwo = 0;

  var rl = require('readline').createInterface({
    input: require('fs').createReadStream('day05.input')
  });

  rl.on('line', function(line) {
    if (day05.isNiceString(line)) {
      niceCount += 1;
    }
    if (day05.isNiceStringPartTwo(line)) {
      niceCountPartTwo += 1;
    }
  });

  rl.on('close', function() {
    callback([niceCount, niceCountPartTwo]);
  });
}

myAnswer(console.log);
