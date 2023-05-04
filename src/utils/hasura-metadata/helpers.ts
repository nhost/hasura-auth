import { logger } from '@/logger';
import { exportMetadata, replaceMetadata } from './api';
import { HasuraMetadataV3, Source } from './types';

/**
 * Extends the original metadata and its source with the new source. Excludes
 * every existing metadata that is related to the `auth` schema and blindly
 * concatenates the new source's tables to the existing ones.
 *
 * @param metadata - The original metadata
 * @param source - The new source
 * @returns The new metadata with the new source
 */
export function extendMetadataSource(
  metadata: HasuraMetadataV3,
  source: Source
): HasuraMetadataV3 {
  return {
    ...metadata,
    sources: metadata.sources.map((originalSource) => {
      if (originalSource.name !== source.name) {
        return originalSource;
      }

      return {
        ...originalSource,
        tables: originalSource.tables
          .filter(({ table }) => table.schema !== 'auth')
          .concat(source.tables),
      };
    }),
  };
}

export const overrideMetadata = async (source: Source) => {
  logger.debug('Exporting metadata...');
  const metadata = await exportMetadata();
  const mergedMetadata = extendMetadataSource(metadata, source);

  logger.debug('Overriding metadata...');
  await replaceMetadata(mergedMetadata);
};
