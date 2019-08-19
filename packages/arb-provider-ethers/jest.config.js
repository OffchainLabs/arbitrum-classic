module.exports = {
    preset: 'ts-jest',
    testEnvironment: 'node',
    reporters: ['default', 'jest-junit'],
    testPathIgnorePatterns: ['/node_modules/', '<rootDir>/dist'],
};
