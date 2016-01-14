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

var myAnswer = function(callback) {
  callback(mine('yzbqklnj', 1, '00000'));
  callback(mine('yzbqklnj', 1, '000000'));
}

myAnswer(console.log);
