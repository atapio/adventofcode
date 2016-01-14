var hasAtLeastThreeVowels = exports.hasAtLeastThreeVowels = function(str) {
  var vowels = 'aeiou'.split("");

  var vowelCount = str.split("").reduce(function(previousValue, currentValue, currentIndex, array) {
    if (vowels.indexOf(currentValue) !== -1) {
      return previousValue + 1;
    }
    return previousValue;
  }, 0);

  return vowelCount >= 3;
};

var hasALetterTwiceInARow = exports.hasALetterTwiceInARow = function(str) {
  var previousValue = null;
  var found = false;

  str.split("").forEach(function(c) {
    if (previousValue === c) {
      found = true;
    }
    previousValue = c;
  });

  return found;
};

var hasIllegalSubStrings = exports.hasIllegalSubStrings = function(str) {
  var illegalStrings = ['ab', 'cd', 'pq', 'xy'];
  var found = false;

  illegalStrings.forEach(function(substring) {
    if (str.indexOf(substring) !== -1) {
      found = true;
    }
  });

  return found;
};

var isNiceString = exports.isNiceString = function(str) {
  return hasAtLeastThreeVowels(str) && hasALetterTwiceInARow(str) && !hasIllegalSubStrings(str);
}
