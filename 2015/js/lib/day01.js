var floor = exports.floor = function(directions, enterBasement) {
  enterBasement = enterBasement || false
  var currentFloor = 0;

  for (i = 0; i < directions.length; i++) {
    if (directions.charAt(i) == '(') {
      currentFloor++;
    } else if (directions.charAt(i) == ')') {
      currentFloor--;
    }

    if (enterBasement && currentFloor < 0) {
      return i + 1;
    }

  }

  return currentFloor;
}
