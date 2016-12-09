var _  = require('underscore');
var fs = require('fs');

try {
  var data = require('./data.json');
  var html = fs.readFileSync('./index.template.html').toString();
  var compiled = _.template(html);
  var result = compiled({ fruits :  data.fruits });
  fs.writeFileSync('result.html', result);
} catch (er) {
  console.error(er);
}
