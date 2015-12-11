var day02 = require("../day02");

exports.testWrappingPaper = function(test){
    test.expect(2);
    test.equal(day02.wrappingPaper("2x3x4"), 58)
    test.equal(day02.wrappingPaper("1x1x10"), 43)
    test.done();
};

