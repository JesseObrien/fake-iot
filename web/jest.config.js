module.exports = {
  moduleFileExtensions: ["js"],
  testEnvironment: "jsdom",
  testMatch: ["**/__tests__/**/*.js?(x)", "**/?(*.)+(spec|test).js?(x)"],
  transformIgnorePatterns: ["<rootDir>/node_modules/"],
  setupFiles: ["<rootDir>/enzyme.config.js"],
  testPathIgnorePatterns: ["\\\\node_modules\\\\"],
  testURL: "http://localhost",
};
