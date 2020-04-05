const name = process.argv[2];
const p = require('../versions.json');
const fs = require('fs');

fs.writeFile(
  'versions.json',
  JSON.stringify(p.filter(v => v !== name)),
  function(err) {
    if (err) {
      return console.error(err);
    }
  }
);
