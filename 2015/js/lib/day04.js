md5 = require('js-md5');

var checkHash = exports.checkHash = function(secret, number, beginningOfHash) {
  if (md5(secret + number).startsWith(beginningOfHash)) {
    return true;
  }
  return false;
};

var mine = exports.mine = function(secret, firstNumber, beginningOfHash) {
  var number = firstNumber;
  while (!checkHash(secret, number, beginningOfHash)) {
    number += 1;
  }
  return number;
}
