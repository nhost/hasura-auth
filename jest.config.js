module.exports = {
    roots: ['./test'],
    globals: {
        server: null,
    },
    verbose: false,
    moduleNameMapper: {
        '^@config$': '<rootDir>/src/config',
        '^@/(.*)$': '<rootDir>/test/src/$1',
        axios: require.resolve('axios'),
    },
    testPathIgnorePatterns: [
        '<rootDir>/dist/',
        '<rootDir>/node_modules/',
        '<rootDir>/data/',
    ],
    setupFilesAfterEnv: ['jest-extended'],
    preset: 'ts-jest',
    testEnvironment: 'node',
};
