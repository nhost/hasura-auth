import { authMetadataSource } from '@/metadata';
import { HasuraMetadataV3, extendMetadataSource } from '@/utils';
import { EMPTY_METADATA } from './empty-metadata';
import { NHOST_PROJECT_METADATA } from './nhost-project-metadata';

describe('metadata operations', () => {
  it('should apply metadata entirely for the auth schema', () => {
    const metadata: HasuraMetadataV3 = { ...EMPTY_METADATA };
    const updatedMetadata = extendMetadataSource(metadata, authMetadataSource);

    expect(updatedMetadata).toMatchSnapshot();
  });

  it('should override existing metadata for the auth schema', () => {
    const metadata: HasuraMetadataV3 = { ...NHOST_PROJECT_METADATA };
    const updatedMetadata = extendMetadataSource(metadata, authMetadataSource);

    expect(updatedMetadata).toMatchSnapshot();
  });
});
