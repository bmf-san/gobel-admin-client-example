module.exports = {
  preset: "@vue/cli-plugin-unit-jest/presets/no-babel",
  moduleNameMapper: {
    "^.+\\.(css|styl|less|sass|scss|png|jpg|ttf|woff|woff2)$":
      "jest-transform-stub",
  },
  transform: {
    "^.+\\.vue$": "vue-jest",
  },
};
