var day03 = require("../day03");

exports.testHouses = function(test){
    test.expect(3);
    test.equal(day03.houses(">"), 2)
    test.equal(day03.houses("^>v<"), 4)
    test.equal(day03.houses("^v^v^v^v^v"), 2)
    test.done();
};

