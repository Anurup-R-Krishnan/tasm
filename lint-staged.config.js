module.exports = {
  // Lint and format Go files in the backend
  'tasm-backend/**/*.go': [
    (filenames) => `bash -c "cd tasm-backend && golangci-lint run ${filenames.map(f => f.replace(process.cwd() + '/tasm-backend/', '')).join(' ')}"`,
    'gofmt -w'
  ],
  // Lint and format Vue/TS/JS files in the frontend
  'tasm-frontend/**/*.{vue,ts,js,cjs,json}': [
    (filenames) => `bash -c "cd tasm-frontend && npm run lint -- ${filenames.map(f => f.replace(process.cwd() + '/tasm-frontend/', '')).join(' ')}"`,
    (filenames) => `bash -c "cd tasm-frontend && npm run format -- ${filenames.map(f => f.replace(process.cwd() + '/tasm-frontend/', '')).join(' ')}"`
  ]
};
