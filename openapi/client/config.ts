import type { ConfigFile } from '@rtk-query/codegen-openapi'

const config: ConfigFile = {
    schemaFile: './_build/pubsubui.yaml',
    apiFile: '@services/config',
    apiImport: 'emptySplitApi',
    outputFile: './index.ts',
    exportName: 'pubsubui',
    hooks: true,
};
export default config;