module.exports = {
  env: {
    commonjs: true,
    es6: true,
    node: true
  },
  plugins: ["prettier"],
  extends: [
    "eslint:recommended",
    "plugin:prettier/recommended" // Enables eslint-plugin-prettier and displays prettier errors as ESLint errors. Make sure this is always the last configuration in the extends array.
  ],
  parserOptions: {
    ecmaVersion: 2018, // Allows for the parsing of modern ECMAScript features
    sourceType: "module" // Allows for the use of imports
  },
  rules: {
    "prettier/prettier": "error",
    "no-unused-vars": "warn"
  },
  overrides: [
    {
      files: ["*.ts", "*.tsx"],
      parser: "@typescript-eslint/parser",
      extends: [
        "eslint:recommended",
        "plugin:@typescript-eslint/recommended",
        "prettier/@typescript-eslint",
        "plugin:prettier/recommended"
      ],
      plugins: ["@typescript-eslint", "prettier"],
      rules: {
        "prettier/prettier": ["error", { singleQuote: true }],
        "@typescript-eslint/no-use-before-define": [
          "error",
          { functions: false }
        ],
        "@typescript-eslint/no-use-before-define": ["warn", { functions: true }]
      }
    }
  ]
};
