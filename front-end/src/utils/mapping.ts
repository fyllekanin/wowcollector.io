import type { HeaderLink, NavigationTree } from '@nuxt/ui-pro/types';
import type {
  AchievementCategory,
  AchievementInformation,
  MountCategory,
  MountInformation,
} from '~/types';

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

export function flatMapAchievements(
  achievementCategories: AchievementCategory[]
): AchievementInformation[] {
  return achievementCategories.reduce((acc, category) => {
    if (category.achievements) acc.push(...category.achievements);
    if (category.categories)
      acc.push(...flatMapAchievements(category.categories));
    return acc;
  }, [] as AchievementInformation[]);
}

export function getRootCategoryNames(
  categories: MountCategory[] | AchievementCategory[]
): string[] {
  return categories.map((category) => category.name);
}

export function getSubCategoryNames(
  categories: MountCategory[] | AchievementCategory[]
): string[] {
  return [
    ...new Set(
      categories.reduce((acc, category) => {
        if (category.categories)
          acc.push(...getRootCategoryNames(category.categories));
        return acc;
      }, [] as string[])
    ),
  ];
}
