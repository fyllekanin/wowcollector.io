import { readdirSync, readFileSync, statSync, writeFileSync } from 'fs';
import { join } from 'path';
import semver from 'semver';

const changelogDir = join(import.meta.url ? new URL('changelog/releases', import.meta.url).pathname : 'changelog/releases');

function getSortedFolders() {
    const folders = readdirSync(changelogDir).filter((folder) => {
        return statSync(join(changelogDir, folder)).isDirectory();
    });

    const sortedFolders = folders
        .sort((a, b) => {
            const versionA = a.split('-')[0];
            const versionB = b.split('-')[0];
            if (versionA === 'latest') return -1;
            if (versionB === 'latest') return -1;
            
            return semver.rcompare(versionA, versionB)
        });

    return sortedFolders;
}

export function getChanges(folderPath) {
    const jsonFiles = [];

    const files = readdirSync(folderPath).filter((file) => {
        return statSync(join(folderPath, file)).isFile() && file.endsWith('.json');
    });

    for (const file of files) {
        const filePath = join(folderPath, file);
        const fileContent = readFileSync(filePath, 'utf-8');
        try {
            const jsonData = JSON.parse(fileContent);
            jsonFiles.push(jsonData);
        } catch (error) {
            console.error(`Error parsing JSON file ${filePath}:`, error);
        }
    }

    return jsonFiles;
}


function getVersionsPopulated(folders) {
    const result = [];

    for (const folder of folders) {
        const [version, timestamp] = folder.split('-');

        result.push({
            version: version,
            timestamp: timestamp,
            changes: getChanges(join(changelogDir, folder))
        })
    }
    return result;
}

function writeTextToFile(text) {
    const filePath = join(import.meta.url ? new URL('server/api/changelog.json.get.ts', import.meta.url).pathname : 'server/api/changelog.json.get.ts');

    writeFileSync(filePath, text, 'utf-8');
    console.log(`File written successfully at ${filePath}`);
}

const changelogContent = `export enum ChangelogType {
  BREAKING = '1',
  FEATURE = '2',
  BUG = '3'
}
  
interface Changelog {
  version: string;
  timestamp: string;
  changes: Change[];
}

interface Change {
  type: ChangelogType;
  content: string;
}

const data: Changelog[] =
${JSON.stringify(getVersionsPopulated(getSortedFolders()), null, 2)} as Changelog[];

export default defineEventHandler({
  handler: () => data,
});`;

writeTextToFile(changelogContent);
