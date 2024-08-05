import type { HeaderLink, NavigationTree } from '@nuxt/ui-pro/types';
import type {
  AchievementCategory,
  AchievementInformation,
  MountCategory,
  MountInformation,
  PetCategory,
  PetInformation,
  ToyCategory,
  ToyInformation,
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

export function flatMapToys(toyCategories: ToyCategory[]): ToyInformation[] {
  return toyCategories.reduce((acc, category) => {
    if (category.toys) acc.push(...category.toys);
    if (category.categories) acc.push(...flatMapToys(category.categories));
    return acc;
  }, [] as ToyInformation[]);
}

export function flatMapPets(petCategories: PetCategory[]): PetInformation[] {
  return petCategories.reduce((acc, category) => {
    if (category.pets) acc.push(...category.pets);
    if (category.categories) acc.push(...flatMapPets(category.categories));
    return acc;
  }, [] as PetInformation[]);
}

export function getRootCategoryNames(
  categories:
    | MountCategory[]
    | AchievementCategory[]
    | ToyCategory[]
    | PetCategory[]
): string[] {
  return categories.map((category) => category.name);
}

export function getSubCategoryNames(
  categories:
    | MountCategory[]
    | AchievementCategory[]
    | ToyCategory[]
    | PetCategory[]
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
