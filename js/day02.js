var wrappingPaper = exports.wrappingPaper = function(box) {
    var edges = box.split('x').map(function(item) {
        return parseInt(item);
    });

    edges = edges.sort(function(a,b) {
        return a-b;
    });

    var sides = [edges[0] * edges[1], edges[0] * edges[2], edges[1] * edges[2]]

    sides = sides.sort(function(a,b) {
        return a-b;
    });

    return 2*sides[0] + 2 * sides[1] + 2 * sides[2] + sides[0];

}

var myAnswer = function(callback) {
    var area = 0;

    var rl = require('readline').createInterface({
        input: require('fs').createReadStream('day02.input')
    });

    rl.on('line', function (line) {
        area += wrappingPaper(line);
    });

    rl.on('close', function () {
        callback(area);
    });
}

myAnswer(console.log);
