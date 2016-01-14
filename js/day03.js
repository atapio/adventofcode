var day03 = require("./lib/day03");

var myAnswer = function(callback) {
  var houseCount = 0;
  var santaAndRobotCount = 0;

  var rl = require('readline').createInterface({
    input: require('fs').createReadStream('day03.input')
  });

  rl.on('line', function(line) {
    houseCount += day03.houses(line);
    santaAndRobotCount += day03.santaAndRobot(line);
  });

  rl.on('close ', function() {
    callback([houseCount, santaAndRobotCount]);
  });
}

myAnswer(console.log);
