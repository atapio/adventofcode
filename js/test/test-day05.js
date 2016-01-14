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

exports.testHasPairOfTwoLetters = function(test) {
  test.expect(7);

  test.ok(day05.hasPairOfTwoLetters("xyxy"));
  test.ok(day05.hasPairOfTwoLetters("aabcdefgaa"));
  test.ok(day05.hasPairOfTwoLetters("qjhvhtzxzqqjkmpb"));
  test.ok(day05.hasPairOfTwoLetters("xxyxx"));
  test.ok(day05.hasPairOfTwoLetters("uurcxstgmygtbstg"));

  test.equal(day05.hasPairOfTwoLetters("aaa"), false);
  test.equal(day05.hasPairOfTwoLetters("ieodomkazucvgmuy"), false);

  test.done();
}

exports.testHasLetterThatRepeatsAfterAChar = function(test) {
  test.expect(8);

  test.ok(day05.hasLetterThatRepeatsAfterAChar("xyx"));
  test.ok(day05.hasLetterThatRepeatsAfterAChar("abcdefeghi"));
  test.ok(day05.hasLetterThatRepeatsAfterAChar("aaa"));
  test.ok(day05.hasLetterThatRepeatsAfterAChar("qjhvhtzxzqqjkmpb"));
  test.ok(day05.hasLetterThatRepeatsAfterAChar("xxyxx"));
  test.ok(day05.hasLetterThatRepeatsAfterAChar("ieodomkazucvgmuy"));

  test.equal(day05.hasLetterThatRepeatsAfterAChar("abc"), false);
  test.equal(day05.hasLetterThatRepeatsAfterAChar("uurcxstgmygtbstg"), false);

  test.done();
};


exports.testIsNiceString = function(test) {
  test.expect(5);

  test.ok(day05.isNiceString("ugknbfddgicrmopn"));
  test.ok(day05.isNiceString("aaa"));

  test.equal(day05.isNiceString("jchzalrnumimnmhp"), false);
  test.equal(day05.isNiceString("haegwjzuvuyypxyu"), false);
  test.equal(day05.isNiceString("dvszwmarrgswjxmb"), false);

  test.done();
}

exports.testIsNiceStringPartTwo = function(test) {
  test.expect(4);

  test.ok(day05.isNiceStringPartTwo("qjhvhtzxzqqjkmpb"));
  test.ok(day05.isNiceStringPartTwo("xxyxx"));

  test.equal(day05.isNiceStringPartTwo("uurcxstgmygtbstg"), false);
  test.equal(day05.isNiceStringPartTwo("ieodomkazucvgmuy"), false);

  test.done();
}
