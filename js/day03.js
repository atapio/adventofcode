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
        key =  loc[0] + '-' + loc[1];
        if (houses[key]) {
            houses[key] += 1;
        } else {
            houses[key] = 1;
        };
        return houses;
    }

    var currentLocation = [0, 0]
        var visitedHouses = {};
    addVisit(visitedHouses, currentLocation);

    for(i=0; i < directions.length; i++) {
        currentLocation = move(directions.charAt(i), currentLocation);
        addVisit(visitedHouses, currentLocation);
    }

    return Object.keys(visitedHouses).length
}

var myAnswer = function(callback) {
    var houseCount = 0;

    var rl = require('readline').createInterface({
        input: require('fs').createReadStream('day03.input')
    });

    rl.on('line', function (line) {
        houseCount += houses(line);
    });

    rl.on('close', function () {
        callback([houseCount]);
    });
}

myAnswer(console.log);
