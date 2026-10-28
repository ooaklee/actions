const assert = require('node:assert/strict');
const { test } = require('node:test');
const { generate } = require('./action-docs.cjs');

const metadata = {
  name: 'Example action',
  description: 'An action for testing documentation generation',
  runs: { using: 'node24' },
  inputs: {
    name: { description: 'The name to greet', required: true },
    repetition: {
      description: 'The number of times to repeat the greeting',
      default: '1',
    },
  },
};
const fixture = ['header', 'description', 'runs', 'inputs', 'outputs']
  .map((name) => {
    const marker = `<!-- action-docs-${name} source="action.yml" -->`;
    return `${marker}\nstale\n${marker}`;
  })
  .join('\n\n');

test('generates metadata deterministically and preserves surrounding prose', () => {
  const source = `Intro\n${fixture}\nUsage\n`;
  const result = generate(source, metadata);
  assert.ok(result.startsWith('Intro\n'));
  assert.ok(result.endsWith('\nUsage\n'));
  assert.match(result, /This action is a `node24` action\./);
  assert.match(
    result,
    /\| `name` \| <p>The name to greet<\/p> \| `true` \| `""` \|/,
  );
  assert.match(
    result,
    /\| `repetition` \| <p>The number of times to repeat the greeting<\/p> \| `false` \| `1` \|/,
  );
  assert.equal(generate(result, metadata), result);
});

test('escapes table values and supports optional inputs and outputs', () => {
  const result = generate(fixture, {
    ...metadata,
    inputs: { optional: { description: 'a|b\n<script>`', default: false } },
    outputs: { result: { description: 'A&B' } },
  });
  assert.match(result, /a&#124;b<br>&lt;script&gt;&#96;/);
  assert.match(result, /\| `false` \| `false` \|/);
  assert.match(result, /## Outputs\n/);
  assert.match(result, /A&amp;B/);
});

test('rejects invalid YAML boolean values and incomplete markers', () => {
  assert.throws(
    () =>
      generate(fixture, {
        ...metadata,
        inputs: { name: { required: 'false' } },
      }),
    /required must be a YAML boolean/,
  );
  assert.throws(() => generate('', metadata), /exactly two header markers/);
  assert.throws(() => generate(fixture, {}), /metadata must include/);
});

test('escapes marker-like metadata and renders empty defaults visibly', () => {
  const result = generate(fixture, {
    ...metadata,
    name: 'A <name>',
    description: '<!-- action-docs-inputs source="action.yml" -->',
    inputs: { optional: { default: '' } },
  });
  assert.match(result, /# A &lt;name&gt;/);
  assert.match(result, /&lt;!-- action-docs-inputs/);
  assert.match(result, /\| `optional` \| <p><\/p> \| `false` \| `""` \|/);
  assert.equal(
    generate(result, {
      ...metadata,
      name: 'A <name>',
      description: '<!-- action-docs-inputs source="action.yml" -->',
      inputs: { optional: { default: '' } },
    }),
    result,
  );
});
