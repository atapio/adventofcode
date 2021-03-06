var day03 = require("../lib/day03");

exports.testHouses = function(test) {
  test.expect(3);
  test.equal(day03.houses(">"), 2);
  test.equal(day03.houses("^>v<"), 4);
  test.equal(day03.houses("^v^v^v^v^v"), 2);
  test.done();
};

exports.testSantaAndRobot = function(test) {
  test.expect(3);
  test.equal(day03.santaAndRobot("^v"), 3);
  test.equal(day03.santaAndRobot("^>v<"), 3);
  test.equal(day03.santaAndRobot("^v^v^v^v^v"), 11);
  test.done();
};
