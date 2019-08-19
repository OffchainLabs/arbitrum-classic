module.exports = {
    plugins: ['prettier'],
    extends: [
        'eslint:recommended',
        'prettier/@typescript-eslint', // Uses eslint-config-prettier to disable ESLint rules from @typescript-eslint/eslint-plugin that would conflict with prettier
        'plugin:prettier/recommended', // Enables eslint-plugin-prettier and displays prettier errors as ESLint errors. Make sure this is always the last configuration in the extends array.
    ],
    parserOptions: {
        ecmaVersion: 2018, // Allows for the parsing of modern ECMAScript features
        sourceType: 'module', // Allows for the use of imports
    },
    rules: {
        'prettier/prettier': 'error',
    },
    overrides: [
        {
            files: ['*.ts', '*.tsx'],
            parser: '@typescript-eslint/parser',
            extends: [
                'eslint:recommended',
                'plugin:@typescript-eslint/recommended',
                'prettier/@typescript-eslint',
                'plugin:prettier/recommended',
            ],
            plugins: ['@typescript-eslint', 'prettier'],
            rules: {
                'prettier/prettier': ['error', { singleQuote: true }],
                '@typescript-eslint/no-use-before-define': ['error', { functions: false }],
                '@typescript-eslint/no-use-before-define': ['warn', { functions: true }],
            },
        },
    ],
};
