module.exports = {
  semi: false,
  trailingComma: 'es5',
  singleQuote: true,
  printWidth: 80,
  tabWidth: 2,
  arrowParens: 'avoid',
  bracketSpacing: true,
  overrides: [
    {
      files: '*.sol',
      options: {
        printWidth: 100,
        tabWidth: 4,
        useTabs: false,
        singleQuote: false,
        bracketSpacing: true,
        explicitTypes: 'always',
      },
    },
  ],
}
