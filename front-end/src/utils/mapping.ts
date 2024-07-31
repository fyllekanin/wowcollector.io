import type { HeaderLink, NavigationTree } from '@nuxt/ui-pro/types';
import type { MountCategory, MountInformation } from '~/types';

export function mapNavigationLinks(links: HeaderLink[]): NavigationTree[] {
  return links.map((link) => ({
    ...link,
    children: link.children ? mapNavigationLinks(link.children) : [],
  }));
}

export function flatMapMounts(
  mountCategories: MountCategory[]
): MountInformation[] {
  return mountCategories.reduce((acc, category) => {
    if (category.mounts) acc.push(...category.mounts);
    if (category.categories) acc.push(...flatMapMounts(category.categories));
    return acc;
  }, [] as MountInformation[]);
}

export function getRootCategoryNames(
  mountCategories: MountCategory[]
): string[] {
  return mountCategories.map((category) => category.name);
}

export function getSubCategoryNames(
  mountCategories: MountCategory[]
): string[] {
  return [
    ...new Set(
      mountCategories.reduce((acc, category) => {
        if (category.categories)
          acc.push(...getRootCategoryNames(category.categories));
        return acc;
      }, [] as string[])
    ),
  ];
}
