var day01 = require("../lib/day01");

exports.testFloor = function(test) {
  test.expect(9);
  test.equal(day01.floor("(())"), 0)
  test.equal(day01.floor("()()"), 0)
  test.equal(day01.floor("((("), 3)
  test.equal(day01.floor("(()(()("), 3)
  test.equal(day01.floor("))((((("), 3)
  test.equal(day01.floor("())"), -1)
  test.equal(day01.floor("))("), -1)
  test.equal(day01.floor(")))"), -3)
  test.equal(day01.floor(")())())"), -3)
  test.done();
};

exports.testEnterBasement = function(test) {
  test.expect(2);
  test.equal(day01.floor(")", true), 1)
  test.equal(day01.floor("()())", true), 5)
  test.done();
}
