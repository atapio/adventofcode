var move = function(direction, currentLocation) {
  switch (direction) {
    case '<':
      return [currentLocation[0] - 1, currentLocation[1]];
      break;
    case '>':
      return [currentLocation[0] + 1, currentLocation[1]];
      break;
    case '^':
      return [currentLocation[0], currentLocation[1] + 1];
      break;
    case 'v':
      return [currentLocation[0], currentLocation[1] - 1];
      break;
  }
};

var houses = exports.houses = function(directions) {
  var addVisit = function(houses, loc) {
    key = loc[0] + '-' + loc[1];
    if (houses[key]) {
      houses[key] += 1;
    } else {
      houses[key] = 1;
    };
    return houses;
  }

  var currentLocation = [0, 0];
  var visitedHouses = {};
  addVisit(visitedHouses, currentLocation);

  for (i = 0; i < directions.length; i++) {
    currentLocation = move(directions.charAt(i), currentLocation);
    addVisit(visitedHouses, currentLocation);
  }

  return Object.keys(visitedHouses).length
}

var santaAndRobot = exports.santaAndRobot = function(directions) {
  var addVisit = function(houses, loc) {
    key = loc[0] + '-' + loc[1];
    if (houses[key]) {
      houses[key] += 1;
    } else {
      houses[key] = 1;
    };
    return houses;
  }

  var santaLocation = [0, 0];
  var roboLocation = [0, 0];
  var visitedHouses = {};

  addVisit(visitedHouses, santaLocation);
  addVisit(visitedHouses, roboLocation);

  for (i = 0; i < directions.length; i++) {
    if (i % 2 == 0) {
      santaLocation = move(directions.charAt(i), santaLocation);
      addVisit(visitedHouses, santaLocation);
    } else {
      roboLocation = move(directions.charAt(i), roboLocation);
      addVisit(visitedHouses, roboLocation);
    }
  }

  return Object.keys(visitedHouses).length
}
