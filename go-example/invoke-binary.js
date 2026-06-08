const process = require('node:process');
const os = require('node:os');
const childProcess = require('node:child_process');
const path = require('node:path');

const PLATFORM = process.platform;
const CPU_ARCH = os.arch();

/**
 * Selects the prebuilt action binary for the current GitHub Actions runner.
 *
 * The template ships Linux binaries because GitHub-hosted JavaScript actions
 * execute on the runner where the workflow job is already running.
 *
 * @returns {'action-amd64' | 'action-arm64'} The binary name under dist/.
 * @throws {Error} When the runner platform or CPU architecture is unsupported.
 */
function chooseBinary() {
  if (PLATFORM !== 'linux') {
    throw new Error('Only linux is supported');
  }

  if (CPU_ARCH !== 'x64' && CPU_ARCH !== 'arm64') {
    throw new Error('Only x64 and arm64 are supported');
  }

  if (CPU_ARCH === 'x64') {
    return 'action-amd64';
  }

  return 'action-arm64';
}

const binary = chooseBinary();
const mainScript = path.join(__dirname, 'dist', binary);

// Forward stdio so logs, annotations, prompts, and failures behave like native action output.
const spawnSyncReturns = childProcess.spawnSync(mainScript, {
  stdio: 'inherit',
});

if (spawnSyncReturns.error) {
  throw spawnSyncReturns.error;
}

if (spawnSyncReturns.signal) {
  console.error(
    `Action binary exited due to signal ${spawnSyncReturns.signal}`,
  );
  process.exit(1);
}

process.exit(spawnSyncReturns.status ?? 1);
