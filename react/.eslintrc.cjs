// .eslintrc.cjs
module.exports = {
  env: {
    browser: true,
    es2021: true,
  },
  extends: ["eslint:recommended", "plugin:react/recommended"],
  parser: "@babel/eslint-parser",
  parserOptions: {
    ecmaVersion: "latest",
    sourceType: "module",
    requireConfigFile: false,
    ecmaFeatures: {
      jsx: true,
    },
    babelOptions: {
      presets: ["@babel/preset-react"],
    },
  },
  plugins: ["react"],
  rules: {
    "react/react-in-jsx-scope": "off",
  },
  settings: {
    react: {
      version: "detect",
    },
  },
};
