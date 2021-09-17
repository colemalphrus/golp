const path = require('path');

module.exports = {
  entry: './static/src/index.js',
  output: {
    path: path.resolve(__dirname, './static/js'),
    filename: 'index.bundle.js',
  },
};