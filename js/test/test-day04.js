var day04 = require("../day04");

exports.testCheckHash = function(test) {
  test.expect(4);
  test.ok(day04.checkHash("abcdef", 609043, '00000'));
  test.equal(day04.checkHash("abcdef", 609042, '00000'), false);
  test.ok(day04.checkHash("pqrstuv", 1048970, '00000'));
  test.equal(day04.checkHash("pqrstuv", 1048969, '00000'), false);
  test.done();
};

exports.testMine = function(test) {
  test.expect(2);
  test.equal(day04.mine("abcdef", 609040, '00000'), 609043);
  test.equal(day04.mine("pqrstuv", 1048960, '00000'), 1048970);
  test.done();
};
