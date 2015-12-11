var parseEdges = function(box) {
    var edges = box.split('x').map(function(item) {
        return parseInt(item);
    });

    edges = edges.sort(function(a,b) {
        return a-b;
    });

    return edges;
}

var wrappingPaper = exports.wrappingPaper = function(box) {
    var edges = parseEdges(box);

    var sides = [edges[0] * edges[1], edges[0] * edges[2], edges[1] * edges[2]]

    sides = sides.sort(function(a,b) {
        return a-b;
    });

    return 2*sides[0] + 2 * sides[1] + 2 * sides[2] + sides[0];

}

var ribbon = exports.ribbon = function(box) {
    var edges = parseEdges(box);

    return 2 * (edges[0] + edges[1]) + edges.reduce(function(previousValue, currentValue, currentIndex, array) {
        return previousValue * currentValue;
    });

}

var myAnswer = function(callback) {
    var requiredPaper = 0;
    var requiredRibbon = 0;

    var rl = require('readline').createInterface({
        input: require('fs').createReadStream('day02.input')
    });

    rl.on('line', function (line) {
        requiredPaper += wrappingPaper(line);
        requiredRibbon += ribbon(line);
    });

    rl.on('close', function () {
        callback([requiredPaper, requiredRibbon]);
    });
}

myAnswer(console.log);
