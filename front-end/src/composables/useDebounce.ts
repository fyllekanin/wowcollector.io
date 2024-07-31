export function useDebounce() {
  let timer: NodeJS.Timeout;

  const debounce = <T extends (...args: any[]) => K, K>(
    fn: T,
    delay: number
  ) => {
    return (...args: Parameters<T>) => {
      clearTimeout(timer);
      timer = setTimeout(() => {
        fn(...args);
      }, delay);
    };
  };

  return {
    debounce,
  };
}
