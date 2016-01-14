var day05 = require("../lib/day05");

exports.testHasAtLeastThreeVowels = function(test) {
  test.expect(7);

  test.ok(day05.hasAtLeastThreeVowels("aei"));
  test.ok(day05.hasAtLeastThreeVowels("xazegov"));
  test.ok(day05.hasAtLeastThreeVowels("aeiouaeiouaeiou"));
  test.ok(day05.hasAtLeastThreeVowels("ugknbfddgicrmopn"));
  test.ok(day05.hasAtLeastThreeVowels("aaa"));

  test.equal(day05.hasAtLeastThreeVowels("dvszwmarrgswjxmb"), false);
  test.equal(day05.hasAtLeastThreeVowels("aa"), false);
  test.done();
};

exports.testHasALetterTwiceInARow = function(test) {
  test.expect(5);

  test.ok(day05.hasALetterTwiceInARow("abcdde"));
  test.ok(day05.hasALetterTwiceInARow("aabbccdd"));
  test.ok(day05.hasALetterTwiceInARow("ugknbfddgicrmopn"));
  test.ok(day05.hasALetterTwiceInARow("aaa"));

  test.equal(day05.hasALetterTwiceInARow("jchzalrnumimnmhp"), false);
  test.done();
}

exports.testHasIllegalSubStrings = function(test) {
  test.expect(7);

  test.ok(day05.hasIllegalSubStrings("ab"));
  test.ok(day05.hasIllegalSubStrings("cd"));
  test.ok(day05.hasIllegalSubStrings("pq"));
  test.ok(day05.hasIllegalSubStrings("xy"));
  test.ok(day05.hasIllegalSubStrings("haegwjzuvuyypxyu"));

  test.equal(day05.hasIllegalSubStrings("ugknbfddgicrmopn"), false);
  test.equal(day05.hasIllegalSubStrings("aaa"), false);
  test.done();
}

exports.testIsNiceString = function(test) {
  test.expect(5);

  test.ok(day05.isNiceString("ugknbfddgicrmopn"));
  test.ok(day05.isNiceString("aaa"));

  test.equal(day05.isNiceString("jchzalrnumimnmhp"), false);
  test.equal(day05.isNiceString("haegwjzuvuyypxyu"), false);
  test.equal(day05.isNiceString("dvszwmarrgswjxmb"), false);

  test.done();
}
