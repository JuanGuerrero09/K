// .eslintrc.cjs
module.exports = {
  env: {
    browser: true,
    es2021: true,
  },
  extends: [
    "eslint:recommended",
    "plugin:react/recommended",
    "plugin:@typescript-eslint/recommended", // reglas de TypeScript
  ],
  parser: "@typescript-eslint/parser", // <- Cambia esto
  parserOptions: {
    ecmaVersion: "latest",
    sourceType: "module",
    ecmaFeatures: {
      jsx: true,
    },
  },
  plugins: ["react", "@typescript-eslint"],
  rules: {
    "react/react-in-jsx-scope": "off",
    // Puedes agregar reglas personalizadas aquí si lo deseas
  },
  settings: {
    react: {
      version: "detect",
    },
  },
};
