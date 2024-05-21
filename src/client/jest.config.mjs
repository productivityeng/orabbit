import nextJest from "next/jest.js";

const createJestConfig = nextJest({
  // Provide the path to your Next.js app to load next.config.js and .env files in your test environment
  dir: "./",
});

// Add any custom config to be passed to Jest
/** @type {import('jest').Config} */
const config = {
  // Add more setup options before each test is run
  // setupFilesAfterEnv: ['<rootDir>/jest.setup.js'],
  moduleNameMapper: {
    '^@/actions(.*)$': '<rootDir>/actions$1',
  },
  testEnvironment: "jest-environment-jsdom",
  collectCoverage: true,
  coverageReporters: ["html", "text", "lcov"],
  collectCoverageFrom: ["app/**/*.{ts,tsx}", "components/**/*.{ts,tsx}"],
  coveragePathIgnorePatterns: [".*layout\\.tsx$"],
  fakeTimers: {
    enableGlobally: true,
  },
};

// createJestConfig is exported this way to ensure that next/jest can load the Next.js config which is async
export default createJestConfig(config);